package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRequest struct {
	Username     string             `bson:"username" json:"username" binding:"required,min=3,max=20"`
	Email        string             `bson:"email" json:"email" binding:"required"`
	Password 	 string             `bson:"password" json:"password" binding:"required"`
	Fullname     string             `bson:"fullname" json:"fullname" binding:"required,min=1,max=20"`
}

type UserSigninRequest struct {
	Username     string             `bson:"username" json:"username" binding:"required,min=3,max=20"`
	Password 	 string             `bson:"password" json:"password" binding:"required"`
}

// User represents a user in the system
type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username" json:"username" binding:"required,min=3,max=20"`
	Email        string             `bson:"email" json:"email" binding:"required"`
	PasswordHash string             `bson:"password_hash" json:"-"`
	Fullname     string             `bson:"fullname" json:"fullname" binding:"required,min=1,max=20"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"` // Timestamp of when the user was created
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"` // Timestamp of when the user was last updated
	IsActive     bool               `bson:"is_active" json:"is_active"`   // Is the user active
}
