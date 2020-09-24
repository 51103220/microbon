package core

import "context"

type Executor interface {
	FunctionName() string
	GetApiVersion() string
	Execute(ctx context.Context) (error, interface{})
}
