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
	sr := NewStationRepository(ConfigureRepository(db, "station"))
	cr := NewCommentRepository(sr)
	s := NewStorage(ur, sr, cr)
	assert.NotNil(t, s)
	assert.NotNil(t, s.userRepository)
	assert.NotNil(t, s.User())
	s.userRepository = nil
	assert.NotNil(t, s.User())
	assert.NotNil(t, s.Station())
	s.stationRepository = nil
	assert.NotNil(t, s.Station())
	assert.NotNil(t, s.Comment())
	s.commentRepository = nil
	assert.NotNil(t, s.Comment())
}
