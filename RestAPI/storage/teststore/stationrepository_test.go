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
		Location:    "156.12 1235.2",
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
		Location:    "156.12 1235.2",
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	u, err := ur.FindByID(id)
	assert.Equal(t, u, station)
}

func TestStationFindByEmail(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	_, err := ur.Create(station)
	assert.Nil(t, err)
	u, err := ur.FindByLocation("156.12 1235.2")
	assert.Equal(t, u, station)
}

func TestStationUpdateByID(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	station.Location = "156.12 1sdfsdf"
	err = ur.UpdateByID(id, station)
	assert.Nil(t, err)
	s, err := ur.FindByID(id)
	assert.Nil(t, err)
	assert.EqualValues(t, s.Location, "156.12 1sdfsdf")
}

func TestStationDeleteByID(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	err = ur.DeleteByID(id)
	assert.Nil(t, err)
	s, err := ur.FindByID(id)
	assert.Nil(t, s)
}

func TestStationRead(t *testing.T) {
	ur := NewStationRepository()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	_, err := ur.Create(station)
	assert.Nil(t, err)
	station = &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	_, err = ur.Create(station)
	assert.Nil(t, err)
	users, err := ur.Read(0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}
