package v1

import (
	"log"
	"testing"
	"time"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	teststorage "github.com/Despenrado/ElCharge/RestAPI/storage/teststore"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestNewCommentService(t *testing.T) {
	ur := teststorage.NewUserRepository()
	sr := teststorage.NewStationRepository()
	cr := teststorage.NewCommentRepository(sr)
	s := teststorage.NewStorage(ur, sr, cr)
	c := NewCommentService(s)
	assert.NotNil(t, c)
}

func testHelperComment() (*CommentService, string, string) {
	ur := teststorage.NewUserRepository()
	sr := teststorage.NewStationRepository()
	cr := teststorage.NewCommentRepository(sr)
	s := teststorage.NewStorage(ur, sr, cr)
	c := NewCommentService(s)
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		},
	}
	if err := station.BeforeCreate(); err != nil {
		log.Fatal(err)
	}
	sid, err := s.Station().Create(station)
	if err != nil {
		log.Fatal(err)
	}
	comment := &models.Comment{
		UserID:   "507f1f77bcf86cd799439011",
		Raiting:  3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		},
	}
	comment.BeforeCreate()
	id, err := cr.Create(sid, comment)
	if err != nil {
		log.Fatal(err)
	}
	return c, sid, id
}

func TestCommentCreate(t *testing.T) {
	cs, _, _ := testHelperComment()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.3",
	}
	sid, err := cs.storage.Station().Create(station)
	comment := &models.Comment{
		UserID:   "507f1f77bcf86cd799439011",
		Raiting:  3,
		Text:     "some text",
		UserName: "test username",
	}
	comm, err := cs.CreateComment(sid, comment)
	assert.Nil(t, err)
	assert.NotNil(t, comm)
	comment.ID = primitive.NewObjectID().Hex()
	comment.Raiting = 5
	comm, err = cs.CreateComment(sid, comment)
	assert.Nil(t, err)
	assert.NotNil(t, comm)
	cs.storage.Station().DeleteByID(sid)
}

func TestCommentFindByID(t *testing.T) {
	cs, sid, id := testHelperComment()
	s, err := cs.FindByID(sid, id)
	assert.Nil(t, err)
	assert.NotEqual(t, s, nil)
	s, err = cs.FindByID(sid, "wrongID")
	assert.NotNil(t, err)
	assert.Nil(t, s)
	cs.storage.Station().DeleteByID(sid)
}

func TestCommentUserName(t *testing.T) {
	cs, sid, _ := testHelperComment()
	s, err := cs.FindByUserName(sid, "test username")
	assert.Nil(t, err)
	assert.NotEqual(t, s, nil)
	s, err = cs.FindByUserName(sid, "wrongID")
	assert.NotNil(t, err)
	assert.Nil(t, s)
	cs.storage.Station().DeleteByID(sid)
}

func TestCommentUpdateByID(t *testing.T) {
	cs, sid, id := testHelperComment()
	comm, err := cs.FindByID(sid, id)
	comm.Raiting = 5
	c, err := cs.UpdateByID(sid, id, comm)
	assert.Nil(t, err)
	assert.EqualValues(t, c.Raiting, 5)
	assert.EqualValues(t, c.Text, "some text")
	assert.NotEqualValues(t, c.Raiting, 3)
	assert.NotEqualValues(t, c.Text, "2@email.com")
	cs.storage.Station().DeleteByID(sid)
}

func TestSCommentDeleteByID(t *testing.T) {
	cs, sid, id := testHelperComment()
	comm, err := cs.FindByID(sid, id)
	assert.Nil(t, err)
	assert.NotNil(t, comm)
	err = cs.DeleteByID(sid, id)
	assert.Nil(t, err)
	comm, err = cs.FindByID(sid, id)
	assert.NotNil(t, err)
	assert.Nil(t, comm)
	cs.storage.Station().DeleteByID(sid)
}

func TestCommentRead(t *testing.T) {
	ur, sid, _ := testHelperComment()
	users, err := ur.Read(sid, 0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}
