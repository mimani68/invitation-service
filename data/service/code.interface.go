package service

import (
	"context"
)

type Code interface {
	GetAllCodes(ctx context.Context) (string, error)
}
