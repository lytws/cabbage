package interactor

import (
	"cbg/entities/eacstoken"
	"cbg/entities/euser"
	"cbg/usecase"
	"cbg/usecase/gateway"
	"context"
	"errors"
)

type Auth struct {
	// repo
	userRepo euser.Repository
	// gateway
	authInfoGw gateway.AuthInfo
	// output port
	oPort usecase.AuthOPort
}

func NewAuth(
	userRepo euser.Repository,
	authInfoGw gateway.AuthInfo,
	oPort usecase.AuthOPort,

) *Auth {
	return &Auth{
		userRepo:   userRepo,
		authInfoGw: authInfoGw,
		oPort:      oPort,
	}
}

func (i *Auth) AuthByPassword(ctx context.Context, username string, password string) error {
	// get user entity
	user, err := i.userRepo.GetOneByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, euser.ErrNotFound) {
			return i.oPort.PutDenial(ctx)
		}
		return err
	}
	// check user password
	if !user.PasswordIs(password) {
		return i.oPort.PutDenial(ctx)
	}
	// output access token
	return i.oPort.PutAccessToken(ctx, eacstoken.Generate(user.Username()).String())
}

func (i *Auth) AuthByAuthCode(ctx context.Context, authCode string) error {
	// get auth info
	authInfoDTO, err := i.authInfoGw.Get(ctx, authCode)
	if err != nil {
		if errors.Is(err, gateway.ErrAuthInfoNotFound) {
			return i.oPort.PutDenial(ctx)
		}
		return err
	}
	// output access token
	return i.oPort.PutAccessToken(ctx, eacstoken.Generate(authInfoDTO.Username).String())
}
