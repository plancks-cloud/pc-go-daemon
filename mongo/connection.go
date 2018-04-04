package mongo

import (
	"fmt"

	"github.com/globalsign/mgo/bson"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/globalsign/mgo"
	log "github.com/sirupsen/logrus"
)

var (
	//Session holds the MongoDB connection
	Session *mgo.Session
)

//Init connects to the mongodb and populates the Session object
func Init() {
	log.Debugln("Dialing Mongodb")
	session, err := mgo.Dial("mongo")
	if err != nil {
		panic(fmt.Sprintf("Could not connect to DB: %s", err))
	}
	log.Infoln("📡  Mongodb connected")
	Session = session
}

//Push saves an object into a collection
func Push(obj interface{}) error {
	err := (GetCollection(obj)).Insert(obj)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error pushing to mongo: %s", err))
	}
	return err
}

//GetCollection returns a collections object named by the parameter
func GetCollection(obj interface{}) *mgo.Collection {
	name := util.GetType(obj)
	return &mgo.Collection{Database: Session.DB(name), Name: name, FullName: name + "." + name}
}

//Upsert saves a document if it exists, otherwise inserts it
func Upsert(obj IDableObjectID) error {
	_, err := (GetCollection(obj)).Upsert(bson.M{"_id": obj.DbID()}, obj)
	return err
}

//UpsertWithID saves a document if it exists, otherwise inserts it
func UpsertWithID(id interface{}, obj interface{}) error {
	_, err := (GetCollection(obj)).Upsert(bson.M{"_id": id}, obj)
	return err
}
