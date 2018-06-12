package api

import (
	"encoding/json"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/controller/db"
	"git.amabanana.com/plancks-cloud/pc-go-daemon/model"
	"net/http"
)

//CheckStatus shows if all is ok
func CheckStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	data := struct {
		Contracts []model.Contract
		Bids      []model.Bid
		Wins      []model.Win
	}{
		db.GetContract(),
		db.GetBid(),
		db.GetWin(),
	}
	json.NewEncoder(w).Encode(&data)

}
