package mcp

import "encoding/json"

type Request struct {
	RequestID string  `json:"request_id"`
	Session   Session `json:"session"`
	Action    Action  `json:"action"`
}

type Session struct {
	ID     string            `json:"id"`
	UserID string            `json:"user_id"`
	Budget int               `json:"budget"`
	Meta   map[string]string `json:"meta"`
}

type Action struct {
	Type     ActionType    `json:"type"`
	Tool     *ToolCall     `json:"tool,omitempty"`
	Resource *ResourceRead `json:"resource, omitempty"`
}

type ActionType string

const (
	ActionTool     ActionType = "tool"
	ActionResource ActionType = "resource"
)

type ToolCall struct {
	Name  string          `json:"name"`
	Input json.RawMessage `json:"input"`
}

type ResourceRead struct {
	URI string `json:"uri"`
}
