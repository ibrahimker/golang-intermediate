package repository

import "time"

type IUser interface {
	Register(username, password string) error
	RegisterWithTimestamp(username, password string, createdAt time.Time) error
}

type User struct{}

func NewUser() IUser {
	return &User{}
}

func (u *User) Register(username, password string) error {
	// ceritanya masukin ke db dan sukses
	return nil
}

func (u *User) RegisterWithTimestamp(username, password string, createdAt time.Time) error {
	// ceritanya masukin ke db dan sukses
	return nil
}
