package services

import (
	"log"

	"github.com/ibrahimker/golang-intermediate/session-11/repository"
)

var userRepo repository.IUser

func init() {
	userRepo = repository.NewUser()
}

func RegisterToDB(username, password string) error {
	if err := userRepo.Register("username", "password"); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
