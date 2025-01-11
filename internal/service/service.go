package service

import(
	"context"
)

type DepositInput struct{
	UUID string
	Amount uint64
}

type WithdrawInput struct{
	UUID string
	Amount uint64
}

type TransferInput struct{
	FromUUID string
	ToUUID string
	Amount uint64
}

type User interface{
	GetBalanceByUUID(ctx context.Context, uuid string) (uint64, error)
	Deposit(ctx context.Context, input DepositInput) error
	Withdraw(ctx context.Context, input WithdrawInput) error
	Transfer(ctx context.Context, input TransferInput) error
}