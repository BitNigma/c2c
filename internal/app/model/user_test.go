package model_test

import (
	"testing"
	"website/internal/app/model"

	"github.com/stretchr/testify/assert"
)

func TestUser_validate(t *testing.T) {

	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},

		{
			name: "emty email",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Email = ""
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().ValidationUser())
			} else {
				assert.Error(t, tc.u().ValidationUser())
			}
		})
	}
}

func TestUser_beforeCreate(t *testing.T) {

	u := model.TestUser(t)
	assert.NoError(t, u.BeforeCreate())
	assert.NotNil(t, u.EncryptedPassword)
}
