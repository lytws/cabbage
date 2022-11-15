package ahttp

import (
	"cbg/usecase"
	"context"
	"net/http"
)

type AuthController struct {
	// input port
	iPort usecase.AuthIPort
	// request framework
	req Request
}

func NewAuthController(iPort usecase.AuthIPort, req Request) *AuthController {
	return &AuthController{
		iPort: iPort,
		req:   req,
	}
}

type AuthWithPasswordForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *AuthController) AuthByPassword(ctx context.Context) {
	// bind json body
	form := AuthWithPasswordForm{}
	if err := c.req.BindJsonBody(ctx, &form); err != nil {
		panic(err)
	}
	// invoke the use case
	if err := c.iPort.AuthByPassword(ctx, form.Username, form.Password); err != nil {
		panic(err)
	}
}

func (c *AuthController) AuthByAuthCode(ctx context.Context) {
	if err := c.iPort.AuthByAuthCode(ctx, c.req.UriParam(ctx, "authCode")); err != nil {
		panic(err)
	}
}

type AuthJsonPresenter struct {
	delivery JsonDelivery
}

func NewAuthJsonPresenter(delivery JsonDelivery) *AuthJsonPresenter {
	return &AuthJsonPresenter{
		delivery: delivery,
	}
}

func (p *AuthJsonPresenter) PutAccessToken(ctx context.Context, accessToken string) error {
	return p.delivery.JSON(ctx, http.StatusOK, map[string]string{"accessToken": accessToken})
}

func (p *AuthJsonPresenter) PutDenial(ctx context.Context) error {
	return p.delivery.JSON(ctx, http.StatusForbidden, "access denied")
}
