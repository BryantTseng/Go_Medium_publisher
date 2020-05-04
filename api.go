package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// APIHandler is a handler for calling api
type APIHandler struct {
	Token string
}
type UserProfile struct {
	Data ProfileDetail
}
type ProfileDetail struct {
	Id       string
	Username string
	Name     string
	Url      string
	ImageUrl string
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
func (a APIHandler) GetUserDetail() {
	api := "https://api.medium.com/v1/me"
	client := &http.Client{}
	req, err := http.NewRequest("GET", api, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", "Bearer "+a.Token)
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode == 401 {
		// token not authorized
		log.Fatalln("Token was invalid.")
	}
	var up UserProfile
	body, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &up)
	fmt.Printf(string(body))
}
