package limit

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/dheeraj-sn/mcp-go/internal/domain/limiter"
)

type CreateTool struct {
	svc limiter.Service
}

type createInput struct {
	Key   string `json:"key"`
	Limit int    `json:"limit"`
}

func NewCreateTool(svc limiter.Service) *CreateTool {
	return &CreateTool{
		svc: svc,
	}
}

func (t *CreateTool) Name() string {
	return "create_limit"
}

func (t *CreateTool) Handle(ctx context.Context, input json.RawMessage) (any, error) {
	var in createInput
	if err := json.Unmarshal(input, &in); err != nil {
		return nil, err
	}

	if in.Key == "" {
		return nil, errors.New("key is required")
	}

	if in.Limit <= 0 {
		return nil, errors.New("limit must be greater than 0")
	}

	return t.svc.Create(ctx, in.Key, in.Limit)
}
