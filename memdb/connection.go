package memdb

import (
	"github.com/hashicorp/go-memdb"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"fmt"
	log "github.com/sirupsen/logrus"
)

var (
	db *memdb.MemDB
)

//Init sets up the database
func Init() {

	// Create the DB schema
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"person": &memdb.TableSchema{
				Name: "person",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "Email"},
					},
				},
			},
		},
	}

	// Create a new data base
	var err error
	db, err = memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}

}

//Push stores an object
func Push(obj interface{}) error {
	name := util.GetType(obj)
	txn := db.Txn(true)
	if err := txn.Insert(name, obj); err != nil {
		log.Errorln(fmt.Sprintf("Error pushing to memdb: %s", err))
		txn.Abort()
		return err
	}
	txn.Commit()
	return nil
}
