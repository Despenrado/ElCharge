package mongostorage

import (
	"log"
	"testing"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/stretchr/testify/assert"
)

func testHelperComent() *CommentRepository {
	db, err := ConnectToDB("mongodb://127.0.0.1:27017", "elCharge")
	if err != nil {
		log.Fatal(err)
	}
	ur := NewStationRepository(ConfigureRepository(db, "station"))
	cr := NewCommentRepository(ur)
	return cr
}

func TestCommentCreate(t *testing.T) {
	cr := testHelperComent()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	sid, err := cr.stationRepository.Create(station)
	ti := models.GetTimeNow()
	comment := &models.Comment{
		Rating:   3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	comment.BeforeCreate()
	id, err := cr.Create(sid, comment)
	assert.Nil(t, err)
	assert.NotEqual(t, id, "")
}

func TestCommentFindByID(t *testing.T) {
	cr := testHelperComent()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	sid, err := cr.stationRepository.Create(station)
	ti := models.GetTimeNow()
	comment := &models.Comment{
		Rating:   3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	comment.BeforeCreate()
	id, err := cr.Create(sid, comment)
	assert.Nil(t, err)
	c, err := cr.FindByID(sid, id)
	assert.Nil(t, err)
	assert.NotEqual(t, &models.Comment{}, c)
}

func TestCommentFindByUserName(t *testing.T) {
	cr := testHelperComent()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	sid, err := cr.stationRepository.Create(station)
	ti := models.GetTimeNow()
	comment := &models.Comment{
		Rating:   3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	comment.BeforeCreate()
	_, err = cr.Create(sid, comment)
	assert.Nil(t, err)
	c, err := cr.FindByUserName(sid, "test username")
	assert.Nil(t, err)
	assert.NotEqual(t, &models.Comment{}, c)
}

func TestCommentUpdateByID(t *testing.T) {
	cr := testHelperComent()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	sid, err := cr.stationRepository.Create(station)
	ti := models.GetTimeNow()
	comment := &models.Comment{
		Rating:   3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	comment.BeforeCreate()
	id, err := cr.Create(sid, comment)
	assert.Nil(t, err)
	comment.Rating = 4
	err = cr.UpdateByID(sid, id, comment)
	assert.Nil(t, err)
	c, err := cr.FindByID(sid, id)
	assert.Nil(t, err)
	assert.NotNil(t, c)
	assert.Equal(t, c.Rating, comment.Rating)
}

func TestCommentDeleteByID(t *testing.T) {
	cr := testHelperComent()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	sid, err := cr.stationRepository.Create(station)
	ti := models.GetTimeNow()
	comment := &models.Comment{
		Rating:   3,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	comment.BeforeCreate()
	id, err := cr.Create(sid, comment)
	assert.Nil(t, err)
	err = cr.DeleteByID(sid, id)
	assert.Nil(t, err)
	c, err := cr.FindByID(sid, id)
	assert.NotNil(t, err)
	assert.Nil(t, c)
}
