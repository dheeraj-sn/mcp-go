package tools

import (
	"fmt"
	"sync"
)

type Registry struct {
	mu    sync.RWMutex
	tools map[string]Tool
}

func NewRegistry() *Registry {
	return &Registry{
		tools: make(map[string]Tool),
	}
}

func (r *Registry) Register(tool Tool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if _, ok := r.tools[tool.Name()]; ok {
		panic(fmt.Sprintf("tool %s already registered", tool.Name()))
	}
	r.tools[tool.Name()] = tool
}

func (r *Registry) Get(name string) (Tool, bool) {
	r.mu.Lock()
	defer r.mu.Unlock()
	tool, ok := r.tools[name]
	return tool, ok

}
