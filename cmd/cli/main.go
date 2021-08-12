package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	linkPortal "../.."
)

const dbFileName = "user_auth.json"
const linkPortalURL = "https://143.110.156.164:5000/users/"

func transformJsonToUserLinks(username string, target interface{}) error {
	resp, err := http.Get(linkPortalURL + username)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	return json.NewDecoder(resp.Body).Decode(target)
}

func main() {
	store, close, err := linkPortal.FileSystemUserStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer close()

	username := store.GetUserCreds().Username
	if username == "" {
		log.Fatal("Missing username in config file! ")
	}

	if len(os.Args) == 1 {
		//if link not provided do a GET request
		// to the API for the users links
		fmt.Println("Getting your most recent links..")
		fmt.Println("...")
		var userLinks []linkPortal.UserLinks
		transformJsonToUserLinks(username, &userLinks)

		for i, link := range userLinks {
			if i == 10 {
				break
			}
			fmt.Println(link.Link)
		}
		return
	}
	cli := linkPortal.NewCLI(os.Args[1], store)

	//cli.AuthenticateUser()
	catPtr := flag.String("category", "default", "category the link falls under")
    linkPtr := flag.String("link", "", "link to be uploaded to the portal")
	
	
	flag.Parse()

	link := *linkPtr
	if link == "" {
		fmt.Println("Missing link! please provide")
		return
	}
	cli.AddNewLink(link, *catPtr)

	fmt.Println("processing your URL! :v)")
}
