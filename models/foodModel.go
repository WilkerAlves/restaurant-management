package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Food struct {
	Id        primitive.ObjectID `bson:"_id"`
	FoodId    string             `json:"food_id"`
	Name      *string            `json:"name" validate:"required, mim=2, max=100"`
	Price     *float64           `json:"price" validate:"required"`
	FoodImage *string            `json:"food_image" validate:"required"`
	MenuId    *string            `json:"menu_id" validate:"required"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
