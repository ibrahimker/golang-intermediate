package main

import (
	"errors"
	"fmt"
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
