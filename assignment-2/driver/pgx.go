package driver

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func NewPostgresConn(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, POSTGRES_URL)
}
