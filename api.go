package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// APIHandler is a handler for calling api
type APIHandler struct {
	Token string
}

// UserProfile struct
type UserProfile struct {
	Data ProfileDetail
}

// ProfileDetail struct
type ProfileDetail struct {
	ID       string
	Username string
	Name     string
	URL      string
	ImageURL string
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

// GetUserDetail use for test if the token is valid
func (a APIHandler) GetUserDetail() ProfileDetail {
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
	body, err := ioutil.ReadAll(resp.Body)
	profile := UserProfile{}
	json.Unmarshal(body, &profile)
	return profile.Data
}
