package gwimpl

import (
	"context"
	"errors"

	"cbg/usecase/gateway"
)

type UserModel = gateway.UserDTO

var ErrUserNotFound = gateway.ErrUserNotFound

type MockUserIODriver interface {
	GetUser(username string) (*UserModel, error)
}

type MockUser struct {
	d MockUserIODriver
}

func NewMockUser(d MockUserIODriver) *MockUser {
	return &MockUser{d: d}
}

func (g *MockUser) Get(ctx context.Context, username string) (gateway.UserDTO, error) {
	userModel, err := g.d.GetUser(username)
	if err != nil {
		if errors.Is(err, ErrUserNotFound) {

			return gateway.UserDTO{}, gateway.ErrUserNotFound
		}
		return gateway.UserDTO{}, err
	}
	return gateway.UserDTO{
		Username: userModel.Username,
		Password: userModel.Password,
	}, nil
}
