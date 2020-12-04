package teststorage

import (
	"testing"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/stretchr/testify/assert"
)

func TestStationCreate(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	assert.NotEqual(t, id, "")
	station.ID = id
	assert.Equal(t, ur.db[id], station)
}

func TestStationFindByID(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	u, err := ur.FindByID(id)
	assert.Equal(t, u, station)
}

func TestStationFindByLocation(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	_, err := ur.Create(station)
	assert.Nil(t, err)
	u, err := ur.FindByLocation(156.12, 1235.2)
	assert.Equal(t, u, station)
}

func TestStationFindByName(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name1",
		OwnerID:     "1234567890123456789012345",
		Latitude:    0.0,
		Longitude:   0.0,
	}
	_, err := ur.Create(station)
	station = &models.Station{
		Description: "testText",
		StationName: "station1 name2",
		OwnerID:     "1234567890123456789012345",
		Latitude:    10.0,
		Longitude:   5.0,
	}
	_, err = ur.Create(station)
	station = &models.Station{
		Description: "testText",
		StationName: "station name3",
		OwnerID:     "1234567890123456789012345",
		Latitude:    40.0,
		Longitude:   5.0,
	}
	_, err = ur.Create(station)
	assert.Nil(t, err)
	stations, err := ur.FindByName("1")
	assert.Equal(t, len(stations), 2)
}

func TestStationFindByDescription(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name1",
		OwnerID:     "1234567890123456789012345",
		Latitude:    0.0,
		Longitude:   0.0,
	}
	_, err := ur.Create(station)
	station = &models.Station{
		Description: "testText",
		StationName: "station1 name2",
		OwnerID:     "1234567890123456789012345",
		Latitude:    10.0,
		Longitude:   5.0,
	}
	_, err = ur.Create(station)
	station = &models.Station{
		Description: "testText",
		StationName: "station name3",
		OwnerID:     "1234567890123456789012345",
		Latitude:    40.0,
		Longitude:   5.0,
	}
	_, err = ur.Create(station)
	assert.Nil(t, err)
	stations, err := ur.FindByDescription("t")
	assert.Equal(t, len(stations), 3)
}

func TestStationFindByRadius(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    0.0,
		Longitude:   0.0,
	}
	_, err := ur.Create(station)
	station = &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    10.0,
		Longitude:   5.0,
	}
	_, err = ur.Create(station)
	station = &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    40.0,
		Longitude:   5.0,
	}
	_, err = ur.Create(station)
	assert.Nil(t, err)
	stations, err := ur.FindInRadius(0, 0, 20, 0, 1000)
	assert.Equal(t, len(stations), 2)
}

func TestStationUpdateByID(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	station.Latitude = 156.12
	err = ur.UpdateByID(id, station, "1234567890123456789012345")
	assert.Nil(t, err)
	s, err := ur.FindByID(id)
	assert.Nil(t, err)
	assert.EqualValues(t, s.Latitude, 156.12)
}

func TestStationDeleteByID(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	err = ur.DeleteByID(id, "1234567890123456789012345")
	assert.Nil(t, err)
	s, err := ur.FindByID(id)
	assert.Nil(t, s)
}

func TestStationRead(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	_, err := ur.Create(station)
	assert.Nil(t, err)
	station = &models.Station{
		Description: "testText",
		StationName: "station name",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	_, err = ur.Create(station)
	assert.Nil(t, err)
	users, err := ur.Read(0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func TestStationUpdateRatingByID(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Latitude:    15.12,
		Longitude:   12.2,
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
	comment = &models.Comment{
		UserID:   "test user id",
		Rating:   5,
		Text:     "some text",
		UserName: "test username",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	_, err = cr.Create(sid, comment)
	assert.Nil(t, err)
	err = ur.UpdateRaitingByID(sid)
	assert.Nil(t, err)
	stat, err := ur.FindByID(sid)
	assert.Nil(t, err)
	assert.Equal(t, stat.Rating, float32(4))
}
