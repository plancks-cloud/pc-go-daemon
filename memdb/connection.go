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
			"Contract": &memdb.TableSchema{
				Name: "Contract",
				Indexes: map[string]*memdb.IndexSchema{
					"_id": &memdb.IndexSchema{
						Name:    "_id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "_id"},
					},
				},
			},
			"Bid": &memdb.TableSchema{
				Name: "Bid",
				Indexes: map[string]*memdb.IndexSchema{
					"contractId": &memdb.IndexSchema{
						Name:    "contractId",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "contractId"},
					},
				},
			},
			"Win": &memdb.TableSchema{
				Name: "Win",
				Indexes: map[string]*memdb.IndexSchema{
					"contractId": &memdb.IndexSchema{
						Name:    "contractId",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "contractId"},
					},
				},
			},
			"Service": &memdb.TableSchema{
				Name: "Service",
				Indexes: map[string]*memdb.IndexSchema{
					"contractId": &memdb.IndexSchema{
						Name:    "contractId",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "contractId"},
					},
				},
			},
			"Wallet": &memdb.TableSchema{
				Name: "Wallet",
				Indexes: map[string]*memdb.IndexSchema{
					"contractId": &memdb.IndexSchema{
						Name:    "contractId",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "contractId"},
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
