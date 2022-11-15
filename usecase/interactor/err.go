package interactor

import (
	"cbg/usecase"
	"context"
)

type Err struct {
	oPort usecase.ErrorOPort
}

func NewErr(oPort usecase.ErrorOPort) *Err {
	return &Err{
		oPort: oPort,
	}
}

func (i *Err) Report(ctx context.Context, err error) {
	// do something lik...
	i.oPort.Put(ctx, err)
}

func (i *Err) ReportInternal(ctx context.Context, m usecase.ErrorReportInternalReqModel) {
	// do something...
	i.oPort.PutInternal(ctx)
}
