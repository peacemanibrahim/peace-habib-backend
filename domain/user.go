package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Firstname    string             `json:"firstname" bson:"firstname" validate:"required,alpha"`
	Lastname     string             `json:"lastname" bson:"lastname" validate:"required,alpha"`
	Email        string             `json:"email" bson:"email" validate:"required,email"`
	ProfileImage string             `json:"image_url" bson:"image_url" validate:"omitempty,url"`
	Password     string             `json:"password" bson:"password" validate:"required,min=6"`
	Location     string             `json:"location" bson:"location" validate:"omitempty"`
	PhoneNumber  string             `json:"phone_number" bson:"phone_number" validate:"omitempty,alphanum"`
	CreatedAt    time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at" bson:"updated_at"`
}
