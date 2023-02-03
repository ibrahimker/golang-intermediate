package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
)

// PgxPoolIface defines a little interface for pgxpool functionality.
// Since in the real implementation we can use pgxpool.Pool,
// this interface exists mostly for testing purpose.
type PgxPoolIface interface {
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
}
