package model

import (
	"time"

	"github.com/google/uuid"
)

// Book struct to describe book object.
type User struct {
	ID         uuid.UUID `db:"id" json:"id" validate:"required,uuid"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
	UserID     uuid.UUID `db:"user_id" json:"user_id" validate:"required,uuid"`
	Email      string    `db:"title" json:"title" validate:"required,lte=255"`
	Password   string    `db:"author" json:"author" validate:"required,lte=255"`
	UserStatus int       `db:"book_status" json:"book_status" validate:"required,len=1"`
	UserAttrs  UserAttrs `db:"book_attrs" json:"book_attrs" validate:"required,dive"`
}

// UserAttrs struct to describe User attributes.
type UserAttrs struct {
	Picture     string `json:"picture"`
	Description string `json:"description"`
	Rating      int    `json:"rating" validate:"min=1,max=10"`
	KYC         bool   `json:"kyc"`
}

func NewUser() *User {
	return &User{}
}
