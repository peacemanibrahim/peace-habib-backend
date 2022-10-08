package domain

import "context"

type Service interface {
	CreateUser(ctx context.Context, user User) (*User, error)
}
