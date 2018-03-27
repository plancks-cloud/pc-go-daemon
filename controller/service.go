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

//GetService returns all services stored in the datastore
func GetService() []model.Service {
	var services []model.Service
	mongo.GetCollection(model.Service{}).Find(nil).All(&services)
	for _, service := range services {
		log.Infoln(fmt.Sprintf("Service: %s", service.ID))
	}
	return services
}

//GetOneService returns a single contract
func GetOneService(id string) (model.Service, error) {
	var service model.Service
	err := mongo.GetCollection(&service).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&service)
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
