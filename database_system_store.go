package linkPortal

import (
	"fmt"

	"gorm.io/gorm"
)

type UserLinks struct {
	gorm.Model
	Username string `json:"Username"`
	Link     string `json:"Link"`
	Category string `json:"Category"`
}

type DatabaseUserStore struct {
	Database *gorm.DB
}

func (s *DatabaseUserStore) GetUserLinks(user string) []UserLinks {
	var userLinks []UserLinks

	s.Database.Where("username = ?", user).Find(&userLinks)

	fmt.Println(userLinks)
	return userLinks
}

func (s *DatabaseUserStore) RecordLink(user string, body UserLinks) {
	s.Database.Create(&UserLinks{Username: user, Link: body.Link, Category: body.Category})
}

func (s *DatabaseUserStore) GetUserCreds() UserCredentials {
	return UserCredentials{"pp", "pp"}
}
