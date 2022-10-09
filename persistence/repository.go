package db

import (
	"context"
	"fmt"

	"github.com/peace-habib-exchange/backend/domain"
)

func (d *PeaceRepository) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	_, err := d.DB.Collection("users").InsertOne(ctx, &user)
	if err != nil {
		fmt.Println(err)
		return &user, err
	}
	return &user, nil
}
