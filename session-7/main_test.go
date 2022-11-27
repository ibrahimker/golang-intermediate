package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLuasKubus(t *testing.T) {
	t.Run("hitung luas jika sisi 4", func(t *testing.T) {
		volume := GetLuasKubus(4)
		require.Equal(t, 64, volume)
	})
}

func TestGetUser(t *testing.T) {
	t.Run("apakah user pertama andi", func(t *testing.T) {
		users := GetUser()
		require.Equal(t, 2, len(users))
		require.Equal(t, "andi", users[0])
	})
}

func TestGetLuasKubus1(t *testing.T) {
	type args struct {
		sisi int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "jika sisi 6",
			args: args{
				sisi: 2,
			},
			want: 8,
		},
		{
			name: "jika sisi 3",
			args: args{
				sisi: 3,
			},
			want: 27,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := GetLuasKubus(tt.args.sisi); got != tt.want {
				t.Errorf("GetLuasKubus() = %v, want %v", got, tt.want)
			}
		})
	}
}
