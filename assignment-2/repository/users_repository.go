package repository

import (
	"context"

	"github.com/ibrahimker/golang-intermediate/assignment-2/entity"
	log "github.com/sirupsen/logrus"
)

type IUser interface {
	Insert(ctx context.Context, user *entity.User) error
	GetByUsernamePassword(ctx context.Context, username, password string) (*entity.User, error)
}

// User is responsible to connect user entity with users table in PostgreSQL.
type User struct {
	pool PgxPoolIface
}

// NewUser creates an instance of User.
func NewUser(pool PgxPoolIface) IUser {
	return &User{pool: pool}
}

// Insert inserts the user into the users table and return the user id.
func (u *User) Insert(ctx context.Context, user *entity.User) error {
	query := "INSERT INTO " +
		"users (username, first_name, last_name, password) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id"

	row := u.pool.QueryRow(ctx, query,
		user.Username,
		user.FirstName,
		user.LastName,
		user.Password,
	)

	if err := row.Scan(&user.ID); err != nil {
		log.Warn("Error when scan rows")
		return err
	}

	return nil
}

// GetByUsernamePassword gets a user from PostgreSQL.
// If there isn't any data, it returns error.
func (u *User) GetByUsernamePassword(ctx context.Context, username, password string) (*entity.User, error) {
	log.Info("Start GetByUsernamePassword", username, password)
	query := "SELECT id,username,first_name,last_name,password " +
		"FROM users WHERE username = $1 AND password = $2"

	row := u.pool.QueryRow(ctx, query, username, password)

	user := entity.User{}
	err := row.Scan(&user.ID, &user.Username, &user.FirstName, &user.LastName, &user.Password)
	if err != nil {
		log.WithError(err).Warn("Error when scan rows")
		return nil, err
	}
	log.Infof("Success retrieve user %+v\n", user)
	return &user, nil
}
