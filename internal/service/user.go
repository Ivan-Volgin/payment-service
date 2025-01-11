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

func (u *UserService) Deposit(ctx context.Context, input DepositInput) error {
	return u.userRepo.Deposit(ctx, input.UUID, input.Amount)
}

func (u *UserService) Withdraw(ctx context.Context, input WithdrawInput) error {
	return u.userRepo.Withdraw(ctx, input.UUID, input.Amount)
}

func (u *UserService) Transfer(ctx context.Context, input TransferInput) error {
	return u.userRepo.Transfer(ctx, input.FromUUID, input.ToUUID, input.Amount)
}