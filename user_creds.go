package linkPortal

import (
	"encoding/json"
	"fmt"
	"io"
)

type UserCredentials struct {
	Username string
	Password string
}

func NewUserCreds(rdr io.Reader) (UserCredentials, error) {
	var userCreds UserCredentials
	err := json.NewDecoder(rdr).Decode(&userCreds)
	if err != nil {
		err = fmt.Errorf("problem parsing user credentials, %v", err)
	}

	return userCreds, err
}
