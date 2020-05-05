package main

import "fmt"

func main() {
	a := APIHandler{}
	a.ReadConfig()
	// test token by getting user Detail
	detail := a.GetUserDetail()
	fmt.Printf("User: %s\nUsername: %s", detail.Name, detail.Username)
}
