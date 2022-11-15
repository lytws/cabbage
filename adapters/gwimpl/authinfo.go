package gwimpl

import (
	"cbg/usecase/gateway"
	"context"
	"errors"
)

type AuthInfoModel = gateway.AuthInfoDTO

var ErrAuthInfoNotFound = gateway.ErrAuthInfoNotFound

type MockAuthInfoIODriver interface {
	GetAuthInfo(ctx context.Context, authCode string) (*AuthInfoModel, error)
}

type MockAuthInfo struct {
	d MockAuthInfoIODriver
}

func NewMockAuthInfo(d MockAuthInfoIODriver) *MockAuthInfo {
	return &MockAuthInfo{
		d: d,
	}
}

func (g *MockAuthInfo) Get(ctx context.Context, authCode string) (gateway.AuthInfoDTO, error) {
	dto, err := g.d.GetAuthInfo(ctx, authCode)
	if err != nil {
		if errors.Is(err, ErrAuthInfoNotFound) {
			return gateway.AuthInfoDTO{}, gateway.ErrAuthInfoNotFound
		}
		return gateway.AuthInfoDTO{}, err
	}
	return gateway.AuthInfoDTO{
		Username: dto.Username,
	}, nil
}
