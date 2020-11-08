package v1

import (
	"testing"

	teststorage "github.com/Despenrado/ElCharge/RestAPI/storage/teststore"
	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	ur := teststorage.NewUserRepository()
	sr := teststorage.NewStationRepository()
	cr := teststorage.NewCommentRepository(sr)
	st := teststorage.NewStorage(ur, sr, cr)
	us := NewUserService(st)
	ss := NewStationService(st)
	cs := NewCommentService(st)
	s := NewService(us, ss, cs)
	assert.NotNil(t, s)
}

func TestUser(t *testing.T) {
	ur := teststorage.NewUserRepository()
	sr := teststorage.NewStationRepository()
	cr := teststorage.NewCommentRepository(sr)
	st := teststorage.NewStorage(ur, sr, cr)
	us := NewUserService(st)
	ss := NewStationService(st)
	cs := NewCommentService(st)
	s := NewService(us, ss, cs)
	assert.NotNil(t, s.User())
}
