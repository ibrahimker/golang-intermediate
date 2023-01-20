package repository_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/ibrahimker/golang-intermediate/assignment-2/entity"
	"github.com/ibrahimker/golang-intermediate/assignment-2/repository"
	"github.com/pashagolub/pgxmock"
	"github.com/stretchr/testify/require"
)

func TestInsert(t *testing.T) {
	query := "INSERT INTO " +
		"users (username, first_name, last_name, password) " +
		"VALUES ($1, $2, $3, $4) " +
		"RETURNING id"

	t.Run("database down", func(t *testing.T) {
		mock, _ := pgxmock.NewConn()
		userRepo := repository.NewUser(mock)
		mock.
			ExpectQuery(regexp.QuoteMeta(query)).
			WillReturnError(errors.New("database down"))

		err := userRepo.Insert(context.Background(), &entity.User{
			ID:        0,
			Username:  "",
			FirstName: "",
			LastName:  "",
			Password:  "",
		})
		require.Equal(t, "database down", err.Error())
	})
}