package teststorage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStore(t *testing.T) {
	ur := NewUserRepository()
	sr := NewStationRepository()
	cr := NewCommentRepository(sr)
	s := NewStorage(ur, sr, cr)
	assert.NotNil(t, s)
	assert.NotNil(t, s.userRepository)
}

func TestUser(t *testing.T) {
	ur := NewUserRepository()
	sr := NewStationRepository()
	cr := NewCommentRepository(sr)
	s := NewStorage(ur, sr, cr)
	assert.NotNil(t, s.User())
	s.userRepository = nil
	assert.NotNil(t, s.User())
}

func TestStation(t *testing.T) {
	ur := NewUserRepository()
	sr := NewStationRepository()
	cr := NewCommentRepository(sr)
	s := NewStorage(ur, sr, cr)
	assert.NotNil(t, s.Station())
	s.userRepository = nil
	assert.NotNil(t, s.Station())
}

func TestComment(t *testing.T) {
	ur := NewUserRepository()
	sr := NewStationRepository()
	cr := NewCommentRepository(sr)
	s := NewStorage(ur, sr, cr)
	assert.NotNil(t, s.Comment())
	s.userRepository = nil
	assert.NotNil(t, s.Comment())
}
