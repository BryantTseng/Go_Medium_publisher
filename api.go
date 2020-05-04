package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// APIHandler is a handler for calling api
type APIHandler struct {
	Token string
}

// ReadConfig reads config from file. etc. medium access token
func (a *APIHandler) ReadConfig() {

	file, err := os.Open("configs/config.json")
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	data, _ := ioutil.ReadAll(file)
	json.Unmarshal(data, &a)

}
