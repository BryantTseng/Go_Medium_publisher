package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	a := APIHandler{}
	a.ReadConfig()
	// test token by getting user Detail
	detail := a.GetUserDetail()
	fmt.Printf("User: %s\nUsername: %s\n", detail.Name, detail.Username)
	//Read arguments, only takes files
	filePaths := os.Args[1:]
	for _, filePath := range filePaths {
		_, err := os.Stat(filePath)
		if err != nil {
			log.Fatal(err)
		}
		a.NewPost(filePath)
	}
}
