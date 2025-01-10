package service

import(
	"payment-service/internal/repo"
	"context"
)

type UserService struct{
	userRepo repo.User
}

func NewUserService(userRepo repo.User) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) GetBalanceByUUID(ctx context.Context, uuid string) (uint64, error) {
	return u.userRepo.GetBalanceByUUID(ctx, uuid)
}

func (u *UserService) Deposit(ctx context.Context, input repo.DepositInput) error {
	return u.userRepo.Deposit(ctx, input)
}

func (u *UserService) Withdraw(ctx context.Context, input repo.WithdrawInput) error {
	return u.userRepo.Withdraw(ctx, input)
}

func (u *UserService) Transfer(ctx context.Context, input repo.TransferInput) error {
	return u.userRepo.Transfer(ctx, input)
}