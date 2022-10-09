package services

import (
	"context"
	"time"

	"github.com/peace-habib-exchange/backend/domain"
)

type PeaceService struct {
	PeaceRepository domain.PeaceRepository
}

func (s *PeaceService) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	user.CreatedAt = time.Now()
	return s.PeaceRepository.CreateUser(ctx, user)
}
