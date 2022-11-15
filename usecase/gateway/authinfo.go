package gateway

import (
	"context"
	"errors"
)

type AuthInfoDTO struct {
	Username string `json:"username"`
}

var (
	ErrAuthInfoNotFound = errors.New("auth info not found")
)

type AuthInfo interface {
	Get(ctx context.Context, authCode string) (AuthInfoDTO, error)
}
