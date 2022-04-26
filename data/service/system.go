package service

import (
	"context"
)

type System interface {
	Ping(ctx context.Context) (string, error)
}
