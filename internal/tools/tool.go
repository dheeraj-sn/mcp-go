package tools

import (
	"context"
	"encoding/json"
)

type Tool interface {
	Name() string
	Handle(ctx context.Context, input json.RawMessage) (any, error)
}
