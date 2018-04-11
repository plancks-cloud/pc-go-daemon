package db

import (
	"context"
	"fmt"
	"strconv"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/mem"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/hashicorp/go-memdb"
	log "github.com/sirupsen/logrus"
)

const serviceTable = "Service"

//CreateService creates a new service
func CreateService(service *model.Service) model.MessageOK {
	err := service.Push()
	if err != nil {
		log.Errorln(fmt.Sprintf("Error saving service: %s", err))
		return model.OkMessage(false, err.Error())
	}
	return model.Ok(true)
}

//ServiceExistsByContractId checks if there is a service for a contract
func ServiceExistsByContractId(id string) bool {
	res, err := mem.GetAllByFieldAndValue(serviceTable, contractId, id)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting services: %s", err))
		//TODO: discuss this... super hard
		return false
	}
	items := iteratorToManyServices(res, err)
	return len(items) == 1
}

//CreateServiceFromWin creates a service instance and saves it to the local database. This service
//is created from a win item
func CreateServiceFromWin(win *model.Win) {

	if ServiceExistsByContractId(win.ContractID) {
		return
	}

	contract, err := GetOneContract(win.ContractID)
	if err != nil {
		log.Errorln(fmt.Sprintf("‼️  Error getting contract of the win: %s", err))
		return
	}
	service := model.Service{
		Name:           contract.ServiceName,
		Image:          contract.ImageAMD64,
		HasWorked:      false,
		EffectiveDate:  util.MakeTimestamp(),
		Network:        "",
		HealthyManaged: false,
		Replicas:       contract.Replicas,
		ContractID:     contract.ID}

	service.RunUntil = service.EffectiveDate + (1000 * contract.SecondsToLive)

	log.Debugln(fmt.Sprintf("Creating service object for contractID: %s", win.ContractID))
	CreateService(&service)

	//Ensure that we check the health soon
	go func() { model.DoorBellHealth <- true }()

}

//GetService returns all services stored in the DataStore
func GetService() (services []model.Service) {
	res, err := mem.GetAll(serviceTable)
	return iteratorToManyServices(res, err)
}

//GetServiceStateResult returns all services stored in the datastore
func GetServiceStateResult() (serviceStateResults []model.ServiceStateResult) {
	services := GetService()
	serviceStates := DockerListRunningServices()

	for _, service := range services {

		contract, _ := GetOneContract(service.ContractID)
		if service.Expired(&contract) {
			continue
		}

		item := model.ServiceStateResult{Service: service, ReplicasLive: 0}
		for _, state := range serviceStates {
			if state.Name == service.Name {
				item.ReplicasLive = state.ReplicasRunning
			}
		}
		serviceStateResults = append(serviceStateResults, item)
	}
	return

}

//ReconServices starts and removes services
func ReconServices() {
	servicesNotYetCreated, servicesToBeDeleted := compareRunningServicesToDB()
	if len(servicesNotYetCreated) > 0 || len(servicesToBeDeleted) > 0 {
		log.Infoln(fmt.Sprintf("❄️  Services to create: %d and services to delete: %d", len(servicesNotYetCreated), len(servicesToBeDeleted)))
	}
	createServices(servicesNotYetCreated)
	deleteServices(servicesToBeDeleted)

}

func compareRunningServicesToDB() (
	servicesNotYetCreated []model.Service,
	servicesToBeDeleted []model.ServiceState) {

	desiredServices := GetService()
	existingServices := DockerListRunningServices()

	for _, service := range desiredServices {
		found := false
		i := 0
		for !found && i < len(existingServices) {
			found = service.Name == existingServices[i].Name
			i++
		}
		if !found {
			contract, err := GetOneContract(service.ContractID)
			if err != nil {
				log.Fatalln(fmt.Sprintf("Error getting contract %s: %s", service.ContractID, err))
			}
			if !service.Expired(&contract) {
				servicesNotYetCreated = append(servicesNotYetCreated, service)
			}
		}
	}

	for _, runningService := range existingServices {
		for _, service := range desiredServices {
			contract, err := GetOneContract(service.ContractID)
			if err != nil {
				log.Fatalln(fmt.Sprintf("Error getting contract %s: %s", service.ContractID, err))
			}
			if service.Name == runningService.Name {
				if service.Expired(&contract) {
					servicesToBeDeleted = append(servicesToBeDeleted, runningService)
				}
			}
		}
	}

	return servicesNotYetCreated, servicesToBeDeleted
}

func createServices(services []model.Service) {
	existingServices := DockerListRunningServices()
	found := false

	for _, service := range services {
		found = false
	SearchRunningServicesLoop:
		for _, runningService := range existingServices {
			if service.Name == runningService.Name {
				found = true
				log.Debugln(fmt.Sprintf("> Will not need to create docker service %s", service.Name))
				break SearchRunningServicesLoop
			}
		}
		log.Debugln(fmt.Sprintf("createServices method.. and for %s was found? %s", service.Name, strconv.FormatBool(found)))

		if !found {
			contract, err := GetOneContract(service.ContractID)
			if err != nil {
				log.Errorln(
					fmt.Sprintf("Error getting contract for service with contractID %s, %s",
						service.ContractID, err))
			}

			//Check again not ancient
			if ExpiredContract(&contract) {
				log.Debugln(fmt.Sprintf("createServices method.. and for %s was ancient", service.Name))
				continue
			}

			log.Infoln(fmt.Sprintf("✅  Creating a service for contractId: %s", service.ContractID))
			createService(&service, &contract)
			go func() { model.DoorBellHealth <- true }() //Ensure that the health check is run soon
		}
	}

}

func createService(service *model.Service, contract *model.Contract) {
	log.Debugln(fmt.Sprintf("createService method!"))

	cli, err := client.NewEnvClient()
	ctx := context.Background()
	if err != nil {
		log.Panicln(fmt.Sprintf("Error getting docker client environment: %s", err))
	}

	replicas := uint64(contract.Replicas)

	spec := swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: contract.ServiceName,
		},
		Mode: swarm.ServiceMode{
			Replicated: &swarm.ReplicatedService{
				Replicas: &replicas,
			},
		},
		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: swarm.ContainerSpec{
				Image: service.Image,
			},
			Resources: &swarm.ResourceRequirements{
				Limits: &swarm.Resources{
					MemoryBytes: int64(contract.RequiredMBMemory * 1024 * 1024),
				},
			},
		},
	}

	_, err = cli.ServiceCreate(
		ctx,
		spec,
		types.ServiceCreateOptions{},
	)

	if err != nil {
		log.Errorln(fmt.Sprintf("Error creating docker service: %s", err))
	}
}

func deleteServices(services []model.ServiceState) {
	cli, err := client.NewEnvClient()
	ctx := context.Background()
	if err != nil {
		log.Panicln(fmt.Sprintf("Error getting docker client environment: %s", err))
	}

	for _, service := range services {
		log.Infoln(fmt.Sprintf("🔥  Removing service: %s", service.Name))
		cli.ServiceRemove(ctx, service.ID)
	}
}

//DeleteServicesByContractID deletes a service
func DeleteServicesByContractID(id string) {
	_, err := mem.Delete(serviceTable, contractId, id)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error deleting wins by contractId: %s", err))
	}

}

func iteratorToManyServices(iterator memdb.ResultIterator, err error) (items []model.Service) {
	if err != nil {
		log.Error(err.Error())
		return nil
	}
	if iterator == nil {
		return items
	}
	more := true
	for more {
		next := iterator.Next()
		if next == nil {
			more = false
			continue
		}
		item := next.(*model.Service)
		items = append(items, *item)
	}
	log.Debugln(fmt.Sprintf("Service iterator counts: %d", len(items)))
	return items

}
