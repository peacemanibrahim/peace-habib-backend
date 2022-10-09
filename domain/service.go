package domain

import "context"

type PeaceService interface {
	CreateUser(ctx context.Context, user User) (*User, error)
}
