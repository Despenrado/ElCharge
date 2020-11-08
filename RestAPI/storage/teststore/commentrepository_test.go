package teststorage

import (
	"testing"
	"time"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/stretchr/testify/assert"
)

func TestCommentCreate(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	sid, err := ur.Create(station)
	cr := NewCommentRepository(ur)
	comment := &models.Comment{
		UserID:   "test user id",
		Raiting:  3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		},
	}
	id, err := cr.Create(sid, comment)
	assert.Nil(t, err)
	assert.NotEqual(t, id, "")
	comment.ID = id
	assert.Equal(t, ur.db[sid].Comments[0], *comment)
}

func TestCommentFindByID(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	sid, err := ur.Create(station)
	cr := NewCommentRepository(ur)
	comment := &models.Comment{
		UserID:   "test user id",
		Raiting:  3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		},
	}
	id, err := cr.Create(sid, comment)
	assert.Nil(t, err)
	c, err := cr.FindByID(sid, id)
	assert.Nil(t, err)
	assert.Equal(t, c, comment)
}

func TestCommentFindByUserName(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	sid, err := ur.Create(station)
	cr := NewCommentRepository(ur)
	comment := &models.Comment{
		UserID:   "test user id",
		Raiting:  3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		},
	}
	_, err = cr.Create(sid, comment)
	assert.Nil(t, err)
	c, err := cr.FindByUserName(sid, "test username")
	assert.Nil(t, err)
	assert.Equal(t, c, comment)
}

func TestCommentUpdateByID(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	sid, err := ur.Create(station)
	cr := NewCommentRepository(ur)
	comment := &models.Comment{
		UserID:   "test user id",
		Raiting:  3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		},
	}
	id, err := cr.Create(sid, comment)
	assert.Nil(t, err)
	comment.Raiting = 4
	err = cr.UpdateByID(sid, id, comment)
	assert.Nil(t, err)
	c, err := cr.FindByID(sid, id)
	assert.Nil(t, err)
	assert.EqualValues(t, c.Raiting, comment.Raiting)
}

func TestCommentDeleteByID(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	sid, err := ur.Create(station)
	cr := NewCommentRepository(ur)
	comment := &models.Comment{
		UserID:   "test user id",
		Raiting:  3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		},
	}
	id, err := cr.Create(sid, comment)
	assert.Nil(t, err)
	err = cr.DeleteByID(sid, id)
	assert.Nil(t, err)
	c, err := ur.FindByID(id)
	assert.NotNil(t, err)
	assert.Nil(t, c)
}
