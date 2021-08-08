package main

import (
	"log"
	"net/http"
	"../.."
)

//const dbFileName = "game.db.json"

func main() {
	// web server for Storage service - LinkPortal
	
	server := linkPortal.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}