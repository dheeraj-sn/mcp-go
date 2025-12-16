package mcp

import "net/http"

type Server interface {
	Handler() http.Handler
}
