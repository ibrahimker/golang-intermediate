package main

import (
	"errors"
	"fmt"
	"log"

	"github.com/ibrahimker/golang-intermediate/session-11/repository"
)

func main() {
	fmt.Println(LuasPersegi(4))
}

func LuasPersegi(sisi int) int {
	return sisi * sisi
}

func Register(username, password string) error {
	if username == "" {
		return errors.New("username tidak boleh kosong")
	}
	if password == "" {
		return errors.New("password tidak boleh kosong")
	}

	// ceritanya masukin ke db

	// kalo sukses return nil
	return nil
}

func RegisterToDB(userRepo repository.IUser, username, password string) error {
	if err := userRepo.Register(username, password); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
