package main

import (
	"fmt"
	"log"
	"os"
	"projects/linkPortal"
)

const dbFileName = "user_auth.json"

func main() {
	store, close, err := linkPortal.FileSystemUserStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	if len(os.Args) == 1 {
		//TODO: 
		//if link not provided do a GET request
		// to the API for the users links
		fmt.Println("Getting your links..")
		return
	} 
	cli := linkPortal.NewCLI(os.Args[1], store)
	cli.AuthenticateUser()
	//TODO: update storage service with users new link
	cli.AddNewLink()

	//TODO: make cli call to POST request to webserver
	fmt.Println("processing your URL! :v)\n")
}
