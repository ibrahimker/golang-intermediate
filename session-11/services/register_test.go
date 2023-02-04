package services_test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/ibrahimker/golang-intermediate/session-11/services"
	mock_repository "github.com/ibrahimker/golang-intermediate/session-11/test/mock/repository"
	"github.com/stretchr/testify/require"
)

func TestRegisterToDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("case error", func(t *testing.T) {
		userRepo := mock_repository.NewMockIUser(ctrl)
		userRepo.EXPECT().Register("username", "").Return(errors.New("error unknown"))

		err := services.RegisterToDB("username", "")

		require.Error(t, err)
	})

	t.Run("case sukses", func(t *testing.T) {
		userRepo := mock_repository.NewMockIUser(ctrl)
		userRepo.EXPECT().Register("username", "password").Return(nil)

		err := services.RegisterToDB("username", "password")

		require.NoError(t, err)
	})
}
