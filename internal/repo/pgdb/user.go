package pgdb

import (
	"account-management-service/pkg/postgres"
	"context"
	"fmt"
	"payment-service/pkg/postgres"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo{
	return &UserRepo{pg}
}

