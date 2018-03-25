package mongo

import (
	"fmt"

	"github.com/globalsign/mgo/bson"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"

	"github.com/globalsign/mgo"
)

var (
	//Session holds the MongoDB connection
	Session *mgo.Session
)

//Init connects to the mongodb and populates the Session object
func Init() {
	session, err := mgo.Dial("mongo")
	if err != nil {
		panic(fmt.Sprintf("Could not connect to DB: %s", err))
	}
	Session = session
}

//Push saves an object into a collection
func Push(obj interface{}) error {
	err := (GetCollection(obj)).Insert(obj)
	if err != nil {
		fmt.Println(fmt.Sprintf("Error pushing to mongo: %s", err))
	}
	return err
}

//Get returns a single instance of an object from the database
func Get(result interface{}, query bson.M) error {
	return GetCollection(util.GetType(result)).Find(query).One(&result)
}

//GetByID returns a single instance of an object from the database
func GetByID(result interface{}, query string) error {
	queryObj := GetCollection(util.GetType(result)).FindId(query)
	return queryObj.One(&result)
}

//GetCollection returns a collections object named by the parameter
func GetCollection(obj interface{}) *mgo.Collection {
	name := util.GetType(obj)
	return &mgo.Collection{Database: Session.DB(name), Name: name, FullName: name + "." + name}
}
