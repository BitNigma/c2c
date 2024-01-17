package model

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User ...
type User struct {
	ID                string
	Email             string
	Password          string
	EncryptedPassword string
}

// Validate user ...
func (u *User) ValidationUser() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.Required, validation.Length(1, 100)),
	)

}

// Check user password before create
func (u *User) BeforeCreate() error {
	if len(u.Password) > 0 {
		enc, err := EcnryptString(u.Password)
		if err != nil {
			return err
		}
		u.EncryptedPassword = enc
	}
	return nil
}

// Encrypt user password
func EcnryptString(pass string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
