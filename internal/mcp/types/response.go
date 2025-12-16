package mcp

type Response struct {
	RequestID string       `json:"request_id"`
	Result    any          `json:"result, omitempty"`
	Error     *Error       `json:"error, omitempty"`
	Meta      ResponseMeta `json:"meta, omitempty"`
}

type ResponseMeta struct {
	Retryable bool   `json:"retryable, omitempty"`
	TraceID   string `json:"trace_id, omitempty"`
}
