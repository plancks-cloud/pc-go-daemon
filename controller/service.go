package controller

import (
	"context"
	"fmt"
	"strconv"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/docker/docker/client"
	"github.com/globalsign/mgo/bson"
	log "github.com/sirupsen/logrus"
)

//CreateService creates a new service
func CreateService(service *model.Service) model.MessageOK {
	err := service.Push()
	if err != nil {
		log.Errorln(fmt.Sprintf("Error saving service: %s", err))
		return model.OkMessage(false, err.Error())
	}
	return model.Ok(true)
}

//CreateServiceFromWin creates a service instance and saves it to the local database. This service
//is created from a win item
func CreateServiceFromWin(win *model.Win) {

	log.Infoln("Creating service from win")
	//Check that does not exist first..
	_, possibleError := GetOneServiceByContract(win.ContractID)
	if possibleError != nil {
		log.Infoln(fmt.Sprintf("Could not find service for contractID: %s", win.ContractID))
	} else {
		return
	}

	contract, err := GetOneContract(win.ContractID)
	if err != nil {
		log.Fatalln("Error getting contract of the win: %s", err)
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

	log.Infoln(fmt.Sprintf("Creating service for contractID: %s", win.ContractID))
	CreateService(&service)
}

//GetService returns all services stored in the datastore
func GetService() []model.Service {
	var services []model.Service
	mongo.GetCollection(model.Service{}).Find(nil).All(&services)
	for _, service := range services {
		log.Infoln(fmt.Sprintf("Service: %s", service.ID))
	}
	return services
}

//GetServiceState returns all services stored in the datastore
func GetServiceState() []model.ServiceState {
	services := GetService()
	var results = []model.ServiceState{}
	for _, element := range services {
		item := model.ServiceState{ID: element.ID, Name: element.Name}
		//TODO
		item.ReplicasRequired = 1
		item.ReplicasRunning = 1
		results = append(results, item)
	}
	return results

}

//GetServiceStateResult returns all services stored in the datastore
func GetServiceStateResult() []model.ServiceStateResult {
	services := GetService()
	serviceStates := DockerListRunningServices()

	var results = []model.ServiceStateResult{}
	for _, service := range services {

		item := model.ServiceStateResult{Service: service, ReplicasLive: 0}
		for _, state := range serviceStates {
			if state.Name == service.Name {
				item.ReplicasLive = state.ReplicasRunning
			}
		}
		results = append(results, item)
	}
	return results

}

//GetOneService returns a single contract
func GetOneService(id string) (model.Service, error) {
	var service model.Service
	err := mongo.GetCollection(&service).Find(bson.M{"_id": id}).One(&service)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting bid: %s", err))
	}
	return service, err
}

//GetOneServiceByContract returns a single contract
func GetOneServiceByContract(contractID string) (model.Service, error) {
	var service model.Service
	err := mongo.GetCollection(&service).Find(bson.M{"contractId": contractID}).One(&service)
	return service, err
}

//UpdateService upserts the given bid
func UpdateService(service *model.Service) error {
	err := service.Upsert()
	return err
}

func reconServices() {
	servicesNotYetCreated, servicesToBeDeleted := compareRunningServicesToDB()
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
	log.Infoln(fmt.Sprintf("createServices method.."))
	existingServices := DockerListRunningServices()
	found := false

	for _, service := range services {
		found = false
	SearchRunningServicesLoop:
		for _, runningService := range existingServices {
			if service.Name == runningService.Name {
				found = true
				log.Infoln(fmt.Sprintf("> Will not need to create docker service %s", service.Name))
				break SearchRunningServicesLoop
			}
		}
		log.Infoln(fmt.Sprintf("createServices method.. and for %s was found? %s", service.Name, strconv.FormatBool(found)))

		if !found {
			contract, err := GetOneContract(service.ContractID)
			if err != nil {
				log.Errorln(
					fmt.Sprintf("Error getting contract for service with contractID %s, %s",
						service.ContractID, err))
			}
			log.Infoln(fmt.Sprintf("GOING TO CREATE SERVICE!!!: %s", service.ID))
			createService(&service, &contract)
		}
	}

}

func createService(service *model.Service, contract *model.Contract) {
	log.Infoln(fmt.Sprintf("createService method!"))

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
			ContainerSpec: &swarm.ContainerSpec{
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
		cli.ServiceRemove(ctx, service.ID)
	}
}
