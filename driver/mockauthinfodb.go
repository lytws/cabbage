package driver

import (
	"cbg/adapters/gwimpl"
	"context"
)

type AuthCode string

type MockAuthInfo struct{}

func NewMockAuthInfo() *MockAuthInfo {
	return &MockAuthInfo{}
}

func (d *MockAuthInfo) GetAuthInfo(ctx context.Context, authCode string) (*gwimpl.AuthInfoModel, error) {
	n, ok := ctx.Value(AuthCode(authCode)).(string)
	if !ok {
		return nil, gwimpl.ErrAuthInfoNotFound
	}
	return &gwimpl.AuthInfoModel{
		Username: n,
	}, nil
}

func (d *MockAuthInfo) ContextWithUsername(ctx context.Context, authCode string, username string) context.Context {
	return context.WithValue(ctx, AuthCode(authCode), username)
}
