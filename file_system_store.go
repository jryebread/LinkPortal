package linkPortal

import (
	"encoding/json"
	"fmt"
	"os"
)

type FileSystemUserStore struct {
	database  *json.Encoder
	userCreds UserCredentials
}

func NewFileSystemUserStore(database *os.File) *FileSystemUserStore {
	database.Seek(0, 0)
	userCreds, _ := NewUserCreds(database)
	return &FileSystemUserStore{
		database:  json.NewEncoder(&tape{database}),
		userCreds: userCreds,
	}
}

func FileSystemUserStoreFromFile(path string) (*FileSystemUserStore, func(), error) {
	db, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		return nil, nil, fmt.Errorf("problem opening %s %v", path, err)
	}

	closeFunc := func() {
		db.Close()
	}

	store := NewFileSystemUserStore(db)

	return store, closeFunc, nil
}

// func (f *FileSystemUserStore) RecordWin(playerName string) {
// 	league := f.league

// 	player := league.Find(playerName)

// 	if player != nil {
// 		player.Wins += 1
// 	} else {
// 		f.league = append(league, Player{playerName, 1})
// 	}

// 	//reset json file with new encoded []Player
// 	f.database.Encode(f.league)
// }

func (f *FileSystemUserStore) GetUserCreds() UserCredentials {
	return f.userCreds
}

func (f *FileSystemUserStore) GetUserLinks() []string {
	// this method should retrieve users links 
	// based on user/password combination
	// from webserver
	var links []string

	//TODO: make webserver GET request to get links for user

	return links
}
