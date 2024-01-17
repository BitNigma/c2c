package store_test

import (
	"testing"
	"website/internal/app/model"
	"website/internal/app/store"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {

	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	u, err := s.User().Create(model.TestUser(t))

	assert.NoError(t, err)
	assert.NotNil(t, u)

}

func TestUserRepositoryFindbyEmail(t *testing.T) {
	s, teardown := store.TestStore(t, databaseURL)
	defer teardown("users")

	email := "test@mail.ru"
	_, err := s.User().FindbyEmail(email)
	assert.Error(t, err)

	u := model.TestUser(t)
	u.Email = email
	s.User().Create(u)

	u, err = s.User().FindbyEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}
