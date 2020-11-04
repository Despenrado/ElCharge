package v1

import (
	"testing"

	teststorage "github.com/Despenrado/ElCharge/RestAPI/storage/teststore"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	ur := teststorage.NewUserRepository()
	st := teststorage.NewStorage(ur)
	us := NewUserService(st)
	s := NewService(us)
	assert.NotNil(t, s)
}

func TestUser(t *testing.T) {
	ur := teststorage.NewUserRepository()
	st := teststorage.NewStorage(ur)
	us := NewUserService(st)
	s := NewService(us)
	assert.NotNil(t, s.User())
}
