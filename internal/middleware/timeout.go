package middleware

import (
	"context"
	"encoding/json"
	"time"

	"github.com/dheeraj-sn/mcp-go/internal/tools"
)

type timeoutTool struct {
	next    tools.Tool
	timeout time.Duration
}

func (t *timeoutTool) Name() string {
	return t.next.Name()
}

func (t *timeoutTool) Handle(ctx context.Context, input json.RawMessage) (any, error) {
	ctx, cancel := context.WithTimeout(ctx, t.timeout)
	defer cancel()
	return t.next.Handle(ctx, input)
}

func WithTimeout(d time.Duration) ToolMiddleware {
	return func(next tools.Tool) tools.Tool {
		return &timeoutTool{
			next:    next,
			timeout: d,
		}
	}
}
