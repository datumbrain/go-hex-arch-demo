package service

import (
	"context"
	"hexashop/internal/core/domain"
	"hexashop/internal/port"
	"log"
)

type User struct {
	userRepo port.UserRepo
}

func NewUser(repo port.UserRepo) *User {
	return &User{userRepo: repo}
}

func (svc *User) CreateUser(ctx context.Context, req *domain.User) error {
	if err := svc.userRepo.Create(ctx, req); err != nil {
		log.Println(ctx, err)
		return err
	}

	return nil
}

func (svc *User) GetUser(ctx context.Context, id int) (*domain.User, error) {
	user, err := svc.userRepo.Get(ctx, id)
	if err != nil {
		log.Println(ctx, err)
		return nil, err
	}

	return user, nil
}

func (svc *User) UpdateUser(ctx context.Context, id int, req *domain.User) error {
	req.ID = id
	if err := svc.userRepo.Update(ctx, req); err != nil {
		log.Println(ctx, err)
		return err
	}

	return nil
}

func (svc *User) DeleteUser(ctx context.Context, id int) error {
	if err := svc.userRepo.Delete(ctx, id); err != nil {
		log.Println(ctx, err)
		return err
	}
	return nil
}
