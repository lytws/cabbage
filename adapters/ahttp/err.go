package ahttp

import (
	"context"
	"net/http"
)

type ErrorJsonPresenter struct {
	delivery JsonDelivery
}

func NewErrorJsonPresenter(delivery JsonDelivery) *ErrorJsonPresenter {
	return &ErrorJsonPresenter{
		delivery: delivery,
	}
}

func (p *ErrorJsonPresenter) Put(ctx context.Context, err error) {
	p.delivery.JSON(ctx, http.StatusBadRequest, map[string]string{
		"err": err.Error(),
	})
}

func (p *ErrorJsonPresenter) PutInternal(ctx context.Context) {
	p.delivery.JSON(ctx, http.StatusBadRequest, map[string]string{
		"err": "internal error",
	})
}
