package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Menu struct {
	Id        primitive.ObjectID `bson:"_id"`
	MenuId    string             `json:"menu_id" validate:"required"`
	Name      string             `json:"name" validate:"required"`
	Category  string             `json:"category" validate:"required"`
	StartDate *time.Time         `json:"start_date"`
	EndDate   *time.Time         `json:"end_date"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
}
