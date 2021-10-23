package base

import "context"

type Command interface {
	Execute(ctx context.Context) (string, error)
}
