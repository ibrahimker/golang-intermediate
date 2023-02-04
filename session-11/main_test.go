package main

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	mock_repository "github.com/ibrahimker/golang-intermediate/session-11/test/mock/repository"
	"github.com/stretchr/testify/require"
)

func TestLuasPersegi(t *testing.T) {
	t.Run("case sisi 4", func(t *testing.T) {
		luas := LuasPersegi(4)
		require.Equal(t, 16, luas)
	})
	t.Run("case sisi 6", func(t *testing.T) {
		luas := LuasPersegi(6)
		require.Equal(t, 36, luas)
	})
}

func TestRegister(t *testing.T) {
	t.Run("case username kosong", func(t *testing.T) {
		err := Register("", "password")
		require.Error(t, err)
	})
	t.Run("case password kosong", func(t *testing.T) {
		err := Register("username", "")
		require.Error(t, err)
	})
	t.Run("case sukses", func(t *testing.T) {
		err := Register("username", "password")
		require.NoError(t, err)
	})
}

func TestRegisterToDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	t.Run("case error", func(t *testing.T) {
		userRepo := mock_repository.NewMockIUser(ctrl)
		userRepo.EXPECT().Register("username", "password").Return(errors.New("error db mati"))

		err := RegisterToDB(userRepo, "username", "password")

		require.Error(t, err)
	})

	t.Run("case sukses", func(t *testing.T) {
		userRepo := mock_repository.NewMockIUser(ctrl)
		userRepo.EXPECT().Register("username", "password").Return(nil)

		err := RegisterToDB(userRepo, "username", "password")

		require.NoError(t, err)
	})
}
