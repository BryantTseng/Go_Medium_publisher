package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// APIHandler is a handler for calling api
type APIHandler struct {
	Token string
	ID    string
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

// postbody struct for POST body
type postbody struct {
	Title         string   `json:"title"`
	ContentFormat string   `json:"contentFormat"`
	Content       string   `json:"content"`
	Tags          []string `json:"tags"`
	PublishStatus string   `json:"publishStatus"`
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
func (a *APIHandler) GetUserDetail() ProfileDetail {
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
	a.ID = profile.Data.ID
	return profile.Data
}

// NewPost Send new post to medium
func (a APIHandler) NewPost(fPath string) {
	api := "https://api.medium.com/v1/users/" + a.ID + "/posts"
	client := &http.Client{}
	req, err := http.NewRequest("POST", api, bytes.NewBuffer(makePostBody(fPath)))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Add("Authorization", "Bearer "+a.Token)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	ss, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(ss))
	defer resp.Body.Close()

}

func makePostBody(fPath string) []byte {
	var b = postbody{}
	filename := strings.Split(fPath, "/")
	filenameS := filename[len(filename)-1]
	f := strings.Split(filenameS, ".")
	b.Title = f[0]
	switch f[1] {
	case "md":
		b.ContentFormat = "markdown"
	case "html":
		b.ContentFormat = "html"
	}
	file, err := os.Open(fPath)
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}
	data, _ := ioutil.ReadAll(file)
	b.Content = string(data)
	b.PublishStatus = "draft"
	js, err := json.Marshal(b)
	fmt.Println(string(js))
	return js
}
