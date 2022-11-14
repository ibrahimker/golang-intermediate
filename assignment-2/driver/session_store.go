package driver

import (
	"log"

	"github.com/antonlindstrom/pgstore"
	"github.com/gorilla/sessions"
)

const (
	SESSION_ID   = "test-id"
	POSTGRES_URL = "postgres://postgresuser:postgrespassword@127.0.0.1:5432/postgres?sslmode=disable"
)

var (
	SESSION_AUTH_KEY       = []byte("my-auth-key-very-secret")
	SESSION_ENCRYPTION_KEY = []byte("my-encryption-key-very-secret123")
)

func NewPostgresStore() *pgstore.PGStore {
	store, err := pgstore.NewPGStore(POSTGRES_URL, SESSION_AUTH_KEY, SESSION_ENCRYPTION_KEY)
	if err != nil {
		log.Fatalln("ERROR", err)
	}

	return store
}

func NewCookieStore() *sessions.CookieStore {
	store := sessions.NewCookieStore(SESSION_AUTH_KEY, SESSION_ENCRYPTION_KEY)
	store.Options.Path = "/"
	store.Options.MaxAge = 86400 * 7
	store.Options.HttpOnly = true

	return store
}
