package limiter

import "context"

type Service interface {
	CreateLimit(ctx context.Context, key string, amount int) error
	GetLimit(ctx context.Context, key string) (int, error)
	DeleteLimit(ctx context.Context, key string) error
}

func NewService() Service {
	return &service{}
}

type service struct {
}

func (s *service) CreateLimit(ctx context.Context, key string, amount int) error {
	return nil
}

func (s *service) GetLimit(ctx context.Context, key string) (int, error) {
	return 0, nil
}

func (s *service) DeleteLimit(ctx context.Context, key string) error {
	return nil
}
