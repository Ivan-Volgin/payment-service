package pgdb

import (
	"context"
	"fmt"
	"payment-service/internal/entity"
	"payment-service/pkg/postgres"
	"payment-service/internal/repo/repoerrs"
	"errors"
	// "log"

	// "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	// "github.com/jackc/pgx/v5/pgconn"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo{
	return &UserRepo{pg}
}

func (u *UserRepo) GetBalanceByUUID(ctx context.Context, uuid string) (entity.User, error) {

	sql, args, _ := u.Builder.Select("*").From("users").Where("id = ?", uuid).ToSql()

	var user entity.User
	err := u.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.UUID,
		&user.Balance,
	)
	if err != nil{
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, repoerrs.ErrNotFound
		}
		return entity.User{}, fmt.Errorf("AccountRepo.GetAccountById - r.Pool.QueryRow: %v", err)
	}
	return user, nil
}