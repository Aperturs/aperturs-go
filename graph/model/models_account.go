package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Account struct {
	ID primitive.ObjectID `bson:"_id" json:"id"`
	Email     string    `json:"email"`
	Role      RoleType  `json:"role"`
	Password  string  `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}