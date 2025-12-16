package middleware

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dheeraj-sn/mcp-go/internal/tools"
)

type authTool struct {
	next tools.Tool
}

func (a *authTool) Name() string {
	return a.next.Name()
}

func (a *authTool) Handle(ctx context.Context, input json.RawMessage) (any, error) {
	userID, ok := ctx.Value("user_id").(string)
	if !ok || userID == "" {
		return nil, errors.New("unauthenticated")
	}
	return a.next.Handle(ctx, input)
}

func WithAuth() ToolMiddleware {
	return func(t tools.Tool) tools.Tool {
		return &authTool{
			next: t,
		}
	}
}
