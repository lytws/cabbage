package gateway

import (
	"context"
	"errors"
)

var ErrUserNotFound = errors.New("user not found")

type UserDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User interface {
	Get(ctx context.Context, username string) (UserDTO, error)
}
