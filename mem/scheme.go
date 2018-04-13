package mem

import "github.com/hashicorp/go-memdb"

var (
	// Create the DB schema
	schema = &memdb.DBSchema{
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
)
