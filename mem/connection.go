package mem

import (
	"fmt"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/util"
	"github.com/hashicorp/go-memdb"
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
					"_id": &memdb.IndexSchema{
						Name:    "_id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "_id"},
					},
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
					"_id": &memdb.IndexSchema{
						Name:    "_id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "_id"},
					},
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
					"_id": &memdb.IndexSchema{
						Name:    "_id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "_id"},
					},
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
					"_id": &memdb.IndexSchema{
						Name:    "_id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "_id"},
					},
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

func GetUniqueById(name string, id string) (interface{}, error) {
	txn := getTransaction(false)
	raw, err := txn.First(name, "_id", id)
	return raw, err

}

func GetAllByFieldAndValue(name string, field string, value string) (memdb.ResultIterator, error) {
	txn := getTransaction(false)
	raw, err := txn.Get(name, field, value)
	return raw, err

}

func GetAll(name string) (memdb.ResultIterator, error) {
	txn := getTransaction(false)
	return txn.Get(name, "_id")

}

func getTransaction(write bool) *memdb.Txn {
	return db.Txn(write)
}

//Push stores an object
func Push(obj interface{}) error {
	name := util.GetType(obj)
	txn := db.Txn(true)
	if err := txn.Insert(name, obj); err != nil {
		log.Errorln(fmt.Sprintf("Error pushing to mem: %s", err))
		txn.Abort()
		return err
	}
	txn.Commit()
	return nil
}

func Delete(name string, field string, id string) (int, error) {
	trx := getTransaction(true)
	return trx.DeleteAll(name, field, id)
}
