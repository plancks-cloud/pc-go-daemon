package util

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
	"time"
)

//Post sends byte payload to an endpoint
func Post(url string, jsonBytes []byte) {
	start := time.Now()
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error doing http post: %s", err))
		panic(err)
	}
	defer resp.Body.Close()

	elapsed := time.Since(start)
	log.Debugln(fmt.Sprintf("⏰  Http post @ %s took %s", url, elapsed))

	ioutil.ReadAll(resp.Body)
}

//Get sends a request to a URL and returns the response
func Get(url string) (*http.Response, error) {
	start := time.Now()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error creating request during remote: %s", err))
		return nil, err
	}
	client := &http.Client{}

	elapsed := time.Since(start)
	log.Debugln(fmt.Sprintf("⏰  Http get @ %s took %s", url, elapsed))

	return client.Do(req)
}

//Options sends a request to a URL using method OPTIONS
func Options(url string) {
	start := time.Now()
	req, err := http.NewRequest("OPTIONS", url, nil)
	if err != nil {
		log.Errorln(fmt.Sprintf("Error creating request during remote: %s", err))
		return
	}
	client := &http.Client{}
	client.Do(req)
	elapsed := time.Since(start)
	log.Debugln(fmt.Sprintf("⏰  Http options @ %s took %s", url, elapsed))

}
