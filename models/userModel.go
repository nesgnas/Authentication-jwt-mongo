package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=1000"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=1000"`
	Email         *string            `json:"email" validate:"email,required"`
	Password      *string            `json:"password" validate:"required,min=6"`
	Token         *string            `json:"token"`
	Phone_number  *string            `json:"phone_number"`
	User_type     *string            `json:"user_type" validate:"required,eq=ADMIN"`
	Refresh_token *string            `json:"refresh_token"`
	CreatedAt     time.Time          `json:"created_at"`
	UpdatedAt     time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
}
