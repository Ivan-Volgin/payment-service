package repo

import(
	"context"

)

type User interface {
	GetBalanceByUUID(ctx context.Context, uuid string) (uint64, error)
	Deposit(ctx context.Context, uuid string, amount uint64) error
	Withdraw(ctx context.Context, uuid string, amount uint64) error
	Transfer(ctx context.Context, fromUuid, toUuid string, amount uint64) error
}