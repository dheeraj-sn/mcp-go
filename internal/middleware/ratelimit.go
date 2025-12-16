package middleware

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dheeraj-sn/mcp-go/internal/tools"
)

type rateLimitTool struct {
	next  tools.Tool
	limit int
}

func (r *rateLimitTool) Name() string {
	return r.next.Name()
}

func (r *rateLimitTool) Handle(ctx context.Context, input json.RawMessage) (any, error) {
	if r.limit <= 0 {
		return nil, errors.New("rate limit exceeded")
	}
	return r.next.Handle(ctx, input)
}

func WithRateLimit(limit int) ToolMiddleware {
	return func(next tools.Tool) tools.Tool {
		return &rateLimitTool{
			next:  next,
			limit: limit,
		}
	}
}
