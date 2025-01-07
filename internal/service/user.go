package service

import(
	"payment-service/internal/repo"
)

type User struct{
	userRepo repo.User
}

func (u *userRepo) 