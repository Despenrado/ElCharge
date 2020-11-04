package teststorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	ur := NewUserRepository()
	s := NewStorage(ur)
	assert.NotNil(t, s)
	assert.NotNil(t, s.userRepository)
}

func TestUser(t *testing.T) {
	ur := NewUserRepository()
	s := NewStorage(ur)
	assert.NotNil(t, s.User())
	s.userRepository = nil
	assert.NotNil(t, s.User())
}
