package main

import (
	"testing"

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
