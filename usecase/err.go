package usecase

import "context"

type ErrorReportInternalReqModel struct {
	Err error
	// ModName string
	// ...
}

type ErrorIPort interface {
	Report(ctx context.Context, err error)
	ReportInternal(ctx context.Context, m ErrorReportInternalReqModel)
}

type ErrorOPort interface {
	Put(ctx context.Context, err error)
	PutInternal(ctx context.Context)
}
