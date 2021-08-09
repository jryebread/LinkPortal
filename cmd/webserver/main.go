package main

import (
	"fmt"
	"log"
	"net/http"

	linkPortal "../.."
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

//const dbFileName = "game.db.json"

func NewDatabaseSystemStore() *linkPortal.DatabaseUserStore {
	db, err := gorm.Open(sqlite.Open("final.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database" + err.Error())
	}
	fmt.Println("setting up DB")
	err2 := db.AutoMigrate(&linkPortal.UserLinks{})
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println("setting up DB, done!")

	return &linkPortal.DatabaseUserStore{
		Database: db,
	}
}

func main() {
	// web server for Storage service - LinkPortal
	// create a new Database store
	dbStore := NewDatabaseSystemStore()

	server := linkPortal.NewLinkServer(dbStore)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}

}
