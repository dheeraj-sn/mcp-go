package limiter

import "context"

type Service interface {
	Create(ctx context.Context, key string, limit int) (int, error)
	Get(ctx context.Context, key string) (int, error)
	Delete(ctx context.Context, key string) error
}

func NewService() Service {
	return &service{}
}

type service struct {
}

func (s *service) Create(ctx context.Context, key string, limit int) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (s *service) Get(ctx context.Context, key string) (int, error) {
	// TODO: Implement
	return 0, nil
}

func (s *service) Delete(ctx context.Context, key string) error {
	// TODO: Implement
	return nil
}
