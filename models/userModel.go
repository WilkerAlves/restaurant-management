package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	UserId       string             `json:"user_id"`
	FirstName    *string            `json:"first_name" validate:"required, mim=2,max=100"`
	LastName     *string            `json:"last_name" validate:"required, mim=2,max=100"`
	Password     *string            `json:"password" validate:"required, mim=6"`
	Email        *string            `json:"email" validate:"required"`
	Avatar       *string            `json:"avatar"`
	Phone        *string            `json:"phone" validate:"required"`
	Token        *string            `json:"token"`
	RefreshToken *string            `json:"refresh_token"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
