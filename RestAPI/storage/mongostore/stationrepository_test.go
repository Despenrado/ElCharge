package mongostorage

import (
	"log"
	"testing"
	"time"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/stretchr/testify/assert"
)

func testHelperStation() *StationRepository {
	db, err := ConnectToDB("mongodb://127.0.0.1:27017", "elCharge")
	if err != nil {
		log.Fatal(err)
	}
	ur := NewStationRepository(ConfigureRepository(db, "station"))
	return ur
}

func TestStationCreate(t *testing.T) {
	ur := testHelperStation()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	assert.NotEqual(t, id, "")
	ur.DeleteByID(id)
}

func TestStationFindByID(t *testing.T) {
	ur := testHelperStation()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	station.ID = id
	s, err := ur.FindByID(id)
	assert.Equal(t, s, station)
	ur.DeleteByID(id)
}

func TestStationFindByEmail(t *testing.T) {
	ur := testHelperStation()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	s, err := ur.FindByLocation("156.12 1235.2")
	station.ID = id
	assert.Nil(t, err)
	assert.NotEqual(t, &models.Station{}, s)
	ur.DeleteByID(id)
}

func TestStationUpdateByID(t *testing.T) {
	ur := testHelperStation()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.2",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
			DeleteAt: time.Now(),
		},
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	station.Location = "156.12 1sdfsdf"
	err = ur.UpdateByID(id, station)
	assert.Nil(t, err)
	station.ID = id
	s, err := ur.FindByID(id)
	assert.EqualValues(t, s.Location, "156.12 1sdfsdf")
	//ur.DeleteByID(id)
}

func TestStationDeleteByID(t *testing.T) {
	ur := testHelperStation()
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
	sr := testHelperStation()
	stations, err := sr.Read(0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, stations)
}
