package linkPortal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type CLI struct {
	inputURL string
	store    PlayerStore
}

const linkPortalURL = "http://localhost:5000/users/"

func NewCLI(input string, store PlayerStore) *CLI {
	return &CLI{
		inputURL: input,
		store:    store,
	}
}

func (cli *CLI) AddNewLink() {
	userCreds := cli.store.GetUserCreds()
	//TODO: make POST to persist link to server for user

	fmt.Println(userCreds)

}

func (cli *CLI) AuthenticateUser() {
	userCreds := cli.store.GetUserCreds()
	fmt.Println(userCreds)
	if userCreds.Username == "" || userCreds.Password == "" {
		fmt.Println("User Credentials not found in file!")
		return
	}
	fmt.Printf("User credentials succesfully parsed from file: %s \n", userCreds)

	//send userCreds to HTTP API POST Request
	postBody, _ := json.Marshal(map[string]string{
		"Username": userCreds.Username,
		"Password": userCreds.Password,
	})
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(linkPortalURL+userCreds.Username,
		"application/json", responseBody)
	//Handle Error
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	//Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)
	log.Printf(sb)

}
