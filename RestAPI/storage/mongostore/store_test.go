package mongostorage

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	db, err := ConnectToDB("mongodb://127.0.0.1:27017", "user")
	if err != nil {
		log.Fatal(err)
	}
	ur := NewUserRepository(ConfigureRepository(db, "user"))
	s := NewStorage(ur)
	assert.NotNil(t, s)
	assert.NotNil(t, s.userRepository)
}

func TestUser(t *testing.T) {
	db, err := ConnectToDB("mongodb://127.0.0.1:27017", "user")
	if err != nil {
		log.Fatal(err)
	}
	ur := NewUserRepository(ConfigureRepository(db, "user"))
	s := NewStorage(ur)
	assert.NotNil(t, s.User())
	s.userRepository = nil
	assert.NotNil(t, s.User())
}
