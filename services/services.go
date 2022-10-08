package services

import (
	"context"
	"time"

	"github.com/peace-habib-exchange/backend/domain"
)

type Service struct {
	Repository domain.Repository
}

func (s *Service) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	user.CreatedAt = time.Now()
	return s.Repository.CreateUser(ctx, user)
}
