package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//Post sends byte payload to an endpoint
func Post(url string, jsonBytes []byte) {

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error doing http post: %s", err))
		panic(err)
	}
	defer resp.Body.Close()

	log.Infoln(fmt.Sprintf("Response status: %s", resp.Status))
	r, _ := ioutil.ReadAll(resp.Body)
	log.Infoln(fmt.Sprintf("Response body: %s", string(r)))

}
