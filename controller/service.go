package controller

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mongo"
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
	contract, err := GetOneContract(win.ID)
	if err != nil {
		log.Fatalln("Error getting contract of the win: %s", err)
	}
	service := model.Service{
		Name:           contract.ServiceName,
		Image:          contract.ImageAMD64,
		HasWorked:      false,
		EffectiveDate:  contract.Timestamp,
		Network:        "",
		HealthyManaged: false,
		Replicas:       contract.Replicas,
		ContractID:     contract.ID}

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

//GetOneService returns a single contract
func GetOneService(id string) (model.Service, error) {
	var service model.Service
	err := mongo.GetCollection(&service).Find(bson.M{"_id": id}).One(&service)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error getting bid: %s", err))
	}
	return service, err
}

//UpdateService upserts the given bid
func UpdateService(service *model.Service) error {
	err := service.Upsert()
	return err
}
