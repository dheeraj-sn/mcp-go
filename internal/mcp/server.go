package mcp

import (
	"encoding/json"
	"net/http"

	"github.com/dheeraj-sn/mcp-go/internal/middleware"
	"github.com/dheeraj-sn/mcp-go/internal/tools"
)

type Server struct {
	registry       *tools.Registry
	toolMiddleware middleware.ToolMiddleware
}

func NewServer(opts ...Option) *Server {
	s := &Server{}
	for _, opt := range opts {
		opt(s)
	}

	if s.registry == nil {
		s.registry = tools.NewRegistry()
	}
	return s
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/mcp", s.handleMCP)
	return mux
}

func (s *Server) handleMCP(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

}
