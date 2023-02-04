package repository

type IUser interface {
	Register(username, password string) error
}

type User struct{}

func NewUser() IUser {
	return &User{}
}

func (u *User) Register(username, password string) error {
	// ceritanya masukin ke db dan sukses
	return nil
}
