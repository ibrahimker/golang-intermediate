package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHello2(t *testing.T) {
	res := GetHello2()
	require.Equal(t, "hello world 2", res)
}
