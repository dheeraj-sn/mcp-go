package tools

type Registry interface {
	Register(Tool)
}

func NewRegistry() Registry {
	return &registry{}
}

type registry struct {
	tools map[string]Tool
}

func (r *registry) Register(tool Tool) {
	r.tools[tool.Name()] = tool
}

func (r *registry) Get(name string) Tool {
	tool, ok := r.tools[name]
	if !ok {
		return nil
	}
	return tool
}

func (r *registry) List() []Tool {
	tools := make([]Tool, 0, len(r.tools))
	for _, tool := range r.tools {
		tools = append(tools, tool)
	}
	return tools
}
