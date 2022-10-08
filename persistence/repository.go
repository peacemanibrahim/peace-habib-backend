package db

import (
	"context"
	"fmt"

	"github.com/peace-habib-exchange/backend/domain"
)

func (d *Repository) CreateUser(ctx context.Context, user domain.User) (*domain.User, error) {
	fmt.Println("I got here now")
	_, err := d.DB.Collection("users").InsertOne(ctx, &user)
	if err != nil {
		fmt.Println(err)
		return &user, err
	}
	return &user, nil
}
