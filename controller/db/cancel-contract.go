package db

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mem"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"github.com/hashicorp/go-memdb"
)

const CancelContractTable = "CancelContract"

//GetCancelContractsByContractID returns all contract cancellations
func GetCancelContractsByContractID(contractID string) (items []model.CancelContract) {
	res, err := mem.GetAllByFieldAndValue(CancelContractTable, contractId, contractID)
	return iteratorToManyCancelContract(res, err)
}

func iteratorToManyCancelContract(res memdb.ResultIterator, err error) (items []model.CancelContract) {
	ch := mem.IteratorToChannel(res, err)
	for i := range ch {
		item := i.(model.CancelContract)
		items = append(items, item)
	}
	return items
}
