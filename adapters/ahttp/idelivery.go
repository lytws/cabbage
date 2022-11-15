package ahttp

import (
	"context"
)

type JsonDelivery interface {
	JSON(ctx context.Context, httpCode int, jsonData any) error
}
