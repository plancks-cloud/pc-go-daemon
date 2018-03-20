package mongo

import (
	"fmt"

	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"

	"github.com/globalsign/mgo"
)

var (
	Session *mgo.Session
)

//Init connects to the mongodb and populates the Session object
func Init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic("Could not connect to DB")
	}
	Session = session
}

//Push saves an object into a collection
func Push(obj interface{}) {
	name := util.GetType(obj)
	err := (GetCollection(obj)).Insert(obj)
	if err != nil {
		panic(fmt.Sprintf("Error inserting into collection %s: %s", name, err))
	}
}

func Get(index string, obj interface{}) {

}

//GetCollection returns a collections object named by the parameter
func GetCollection(obj interface{}) *mgo.Collection {
	name := util.GetType(obj)
	return &mgo.Collection{Database: Session.DB(name), Name: name, FullName: name}
}
