package ahttp

import "context"

type Request interface {
	// Header(ctx context.Context, key string) string
	UriParam(ctx context.Context, key string) string
	// QueryParam(ctx context.Context, key string) string
	BindJsonBody(ctx context.Context, modelPointer any) error
}
