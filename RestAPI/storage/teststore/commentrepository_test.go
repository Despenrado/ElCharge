package teststorage

import (
	"testing"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/stretchr/testify/assert"
)

func TestCommentCreate(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	sid, err := ur.Create(station)
	cr := NewCommentRepository(ur)
	ti := models.GetTimeNow()
	comment := &models.Comment{
		UserID:   "test user id",
		Rating:   3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
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
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	sid, err := ur.Create(station)
	cr := NewCommentRepository(ur)
	ti := models.GetTimeNow()
	comment := &models.Comment{
		UserID:   "test user id",
		Rating:   3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
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
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	sid, err := ur.Create(station)
	cr := NewCommentRepository(ur)
	ti := models.GetTimeNow()
	comment := &models.Comment{
		UserID:   "test user id",
		Rating:   3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
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
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	sid, err := ur.Create(station)
	cr := NewCommentRepository(ur)
	ti := models.GetTimeNow()
	comment := &models.Comment{
		UserID:   "test user id",
		Rating:   3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	id, err := cr.Create(sid, comment)
	assert.Nil(t, err)
	comment.Rating = 4
	err = cr.UpdateByID(sid, id, comment)
	assert.Nil(t, err)
	c, err := cr.FindByID(sid, id)
	assert.Nil(t, err)
	assert.EqualValues(t, c.Rating, comment.Rating)
}

func TestCommentDeleteByID(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	sid, err := ur.Create(station)
	cr := NewCommentRepository(ur)
	ti := models.GetTimeNow()
	comment := &models.Comment{
		UserID:   "test user id",
		Rating:   3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
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
