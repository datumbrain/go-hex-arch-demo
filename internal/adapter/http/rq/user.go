package rq

import "hexashop/internal/core/domain"

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

func DomainUser(u User) domain.User {
	return domain.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		IsAdmin:  u.IsAdmin,
	}
}
