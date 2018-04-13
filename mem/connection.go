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
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
			"Bid": &memdb.TableSchema{
				Name: "Bid",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"contractId": &memdb.IndexSchema{
						Name:    "contractId",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ContractID"},
					},
				},
			},
			"Win": &memdb.TableSchema{
				Name: "Win",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"contractId": &memdb.IndexSchema{
						Name:    "contractId",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ContractID"},
					},
				},
			},
			"Service": &memdb.TableSchema{
				Name: "Service",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
					"contractId": &memdb.IndexSchema{
						Name:    "contractId",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "ContractID"},
					},
				},
			},
			"Wallet": &memdb.TableSchema{
				Name: "Wallet",
				Indexes: map[string]*memdb.IndexSchema{
					"id": &memdb.IndexSchema{
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.StringFieldIndex{Field: "ID"},
					},
				},
			},
		},
	}

	// Create a new data base
	dbConnect, err := memdb.NewMemDB(schema)
	if err != nil {
		panic(err)
	}
	db = dbConnect

}

func GetUniqueById(name string, id string) (interface{}, error) {
	txn := getTransaction(false)
	txn.Abort()
	raw, err := txn.First(name, "id", id)
	return raw, err

}

func GetAllByFieldAndValue(name string, field string, value string) (memdb.ResultIterator, error) {
	txn := getTransaction(false)
	txn.Abort()
	raw, err := txn.Get(name, field, value)
	return raw, err

}

func GetAll(name string) (memdb.ResultIterator, error) {
	txn := getTransaction(false)
	txn.Abort()
	return txn.Get(name, "id")

}

func getTransaction(write bool) *memdb.Txn {
	return db.Txn(write)
}

//Push stores an object
func Push(obj interface{}) error {
	name := util.GetType(obj)
	log.Debugln(fmt.Sprintf("Trying to insert a: %s", name))
	txn := db.Txn(true)
	txn.Abort()
	if err := txn.Insert(name, obj); err != nil {
		log.Errorln(fmt.Sprintf("Error pushing to mem: %s", err))
		txn.Abort()
		return err
	}
	txn.Commit()
	return nil
}

func Delete(name string, field string, id string) (int, error) {
	txn := getTransaction(true)
	i, err := txn.DeleteAll(name, field, id)
	if err != nil {
		txn.Abort()
	}
	return i, err
}
