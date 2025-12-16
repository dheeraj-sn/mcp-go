package mcp

import (
	"github.com/dheeraj-sn/mcp-go/internal/middleware"
	"github.com/dheeraj-sn/mcp-go/internal/tools"
)

type Option func(*Server)

func WithToolRegistry(registry *tools.Registry) Option {
	return func(s *Server) {
		s.registry = registry
	}
}

func WithToolMiddleware(mw middleware.ToolMiddleware) Option {
	return func(s *Server) {
		s.toolMiddleware = mw
	}
}
