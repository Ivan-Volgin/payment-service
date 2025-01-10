package postgres

import(
	// "fmt"
	// "log"
	"time"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx"
	"github.com/Masterminds/squirrel"
)

const(
	defaultMaxPoolSize = 1
	defaultConnAttempts = 10
	defaultConnTimeout = time.Second
)

type PgxPool interface{
	Close()
	Acquire(ctx context.Context) (*pgxpool.Conn, error)
	Exec(ctx context.Context, sql string, args ...any) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRaw(ctx context.Context, sql string, args ...any) pgx.Row
	// SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Begin(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames string, rowSrc pgx.CopyFromSource) (int64, error)
	Ping(ctx context.Context) error
}

type Postgres struct {
	maxPoolSize int
	connAttempts int
	connTimeout int

	Builder squirrel.StatementBuilderType
	Pool PgxPool
}