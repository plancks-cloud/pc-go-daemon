package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

//Post comment
func Post(url string, jsonBytes []byte) {
	log.Infoln(fmt.Sprintf("Posting to: %s", url))
	log.Infoln(fmt.Sprintln(string(jsonBytes)))

}

//Post2 sends byte payload to an endpoint
func Post2(url string, jsonBytes []byte) {

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

//Get sends a request to a URL and returns the response
func Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error creating request during sync: %s", err))
		return nil, err
	}
	client := &http.Client{}

	return client.Do(req)
}
