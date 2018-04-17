package db

import (
	"git.amabanana.com/plancks-cloud/pc-go-daemon/mem"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"github.com/hashicorp/go-memdb"
	"fmt"
	log "github.com/sirupsen/logrus"
)

const CancelContractTable = "CancelContract"

//CreateCancelContract creates a new cancel contract
func CreateCancelContract(item *model.CancelContract) model.MessageOK {
	err := item.Push()
	if err != nil {
		log.Errorln(fmt.Sprintf("Error saving cancel contract: %s", err))
		return model.OkMessage(false, err.Error())
	}
	log.Infoln(fmt.Sprintf("❤️  Cancel Contract created ID: %s", item.ID))
	return model.Ok(true)

}

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
