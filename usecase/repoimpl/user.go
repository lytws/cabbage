package repoimpl

import (
	"cbg/entities/euser"
	"cbg/usecase/gateway"
	"context"
	"errors"
)

type User struct {
	userGw gateway.User
}

func NewUser(userGw gateway.User) *User {
	return &User{
		userGw: userGw,
	}
}

func (r *User) GetOneByUsername(ctx context.Context, username string) (*euser.Entity, error) {
	dto, err := r.userGw.Get(ctx, username)
	if err != nil {
		if errors.Is(err, gateway.ErrUserNotFound) {
			return nil, euser.ErrNotFound
		}
		return nil, err
	}
	return euser.Parse(euser.DTO{
		Username: dto.Username,
		Password: dto.Password,
	}), nil
}
