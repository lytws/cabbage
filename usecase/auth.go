package usecase

import "context"

type AuthIPort interface {
	AuthByPassword(ctx context.Context, username string, password string) error
	AuthByAuthCode(ctx context.Context, authCode string) error
}

type AuthOPort interface {
	PutAccessToken(ctx context.Context, accessToken string) error
	PutDenial(ctx context.Context) error
}
