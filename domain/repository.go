package domain

import "context"

type PeaceRepository interface {
	CreateUser(ctx context.Context, user User) (*User, error)
}
