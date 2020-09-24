package core

import "context"

type GBonRequest interface {
	RegisterAs() string
	GetApiVersion() string
	Process(ctx context.Context) (error, interface{})
}
