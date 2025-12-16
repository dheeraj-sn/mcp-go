package middleware

import "github.com/dheeraj-sn/mcp-go/internal/tools"

type ToolMiddleware func(tools.Tool) tools.Tool

func Chain(middlewares ...ToolMiddleware) ToolMiddleware {
	return func(t tools.Tool) tools.Tool {
		wrapped := t
		for i := len(middlewares) - 1; i >= 0; i-- {
			wrapped = middlewares[i](wrapped)
		}
		return wrapped
	}
}
