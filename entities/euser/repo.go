package euser

import (
	"context"
	"errors"
)

var (
	ErrNotFound = errors.New("user not found")
)

type Repository interface {
	GetOneByUsername(ctx context.Context, username string) (*Entity, error)
}
