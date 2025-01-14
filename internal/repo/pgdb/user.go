package pgdb

import (
	"context"
	"fmt"
	"payment-service/internal/entity"
	"payment-service/pkg/postgres"
	"payment-service/internal/repo/repoerrs"
	"errors"
	// "log"

	"github.com/Masterminds/squirrel"
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

	sql, args, _ := u.Builder.Select("*").From("users").Where("uuid = ?", uuid).ToSql()

	var user entity.User
	err := u.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.UUID,
		&user.Balance,
	)
	if err != nil{
		if errors.Is(err, pgx.ErrNoRows) {
			return entity.User{}, repoerrs.ErrNotFound
		}
		return entity.User{}, fmt.Errorf("UserRepo.GetAccountById - r.Pool.QueryRow: %v", err)
	}
	return user, nil
}

func (u *UserRepo) Deposit(ctx context.Context, uuid string, amount uint64) error {
	tx, err := u.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("UserRepo.Deposit - r.Pool.Begin: %v", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	sql, args, _ := u.Builder.
		Update("users").
		Set("balance", squirrel.Expr("balance + ?", amount)).
		Where("uuid = ?", uuid).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo.Deposit - tx.Exec: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("UserRepo.Deposit - tx.Commit: %v", err)
	}
	
	return nil
}

func(u *UserRepo) Withdraw(ctx context.Context, uuid string, amount uint64) error{
	tx, err := u.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("UserRepo.Withdraw - r.Pool.Begin: %v", err)
	}
	defer func() { _ = tx.Rollback(ctx) }()

	sql, args, _ := u.Builder.
		Update("users").
		Set("balance", squirrel.Expr("balance - ?", amount)).
		Where("uuid = ?", uuid).
		ToSql()

	_, err = tx.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UserRepo.Withdraw - tx.Exec: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		return fmt.Errorf("UserRepo.Withdraw - tx.Commit: %v", err)
	}
	
	return nil
}