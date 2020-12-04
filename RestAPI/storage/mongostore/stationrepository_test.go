package mongostorage

import (
	"log"
	"testing"

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
		OwnerID:     "1234567890123456789012345",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	assert.NotEqual(t, id, "")
	ur.DeleteByID(id, "1234567890123456789012345")
}

func TestStationFindByID(t *testing.T) {
	ur := testHelperStation()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	station.ID = id
	s, err := ur.FindByID(id)
	assert.Equal(t, s, station)
	ur.DeleteByID(id, "1234567890123456789012345")
}

func TestStationFindByName(t *testing.T) {
	ur := testHelperStation()
	// station := &models.Station{
	// 	Description: "testText",
	// 	StationName: "station name1",
	// 	Latitude:    0.0,
	// 	Longitude:   0.0,
	// }
	// _, err := ur.Create(station)
	// station = &models.Station{
	// 	Description: "testText",
	// 	StationName: "station1 name2",
	// 	Latitude:    10.0,
	// 	Longitude:   5.0,
	// }
	// _, err = ur.Create(station)
	// station = &models.Station{
	// 	Description: "testText",
	// 	StationName: "station name3",
	// 	Latitude:    40.0,
	// 	Longitude:   5.0,
	// }
	// _, err = ur.Create(station)
	// assert.Nil(t, err)
	stations, err := ur.FindByName("n")
	assert.Nil(t, err)
	assert.NotEqual(t, len(stations), 0)
}

func TestStationFindByDescription(t *testing.T) {
	ur := testHelperStation()
	// station := &models.Station{
	// 	Description: "testText",
	// 	StationName: "station name1",
	// 	Latitude:    0.0,
	// 	Longitude:   0.0,
	// }
	// _, err := ur.Create(station)
	// station = &models.Station{
	// 	Description: "testText",
	// 	StationName: "station1 name2",
	// 	Latitude:    10.0,
	// 	Longitude:   5.0,
	// }
	// _, err = ur.Create(station)
	// station = &models.Station{
	// 	Description: "testText",
	// 	StationName: "station name3",
	// 	Latitude:    40.0,
	// 	Longitude:   5.0,
	// }
	// _, err = ur.Create(station)
	// assert.Nil(t, err)
	stations, err := ur.FindByDescription("t")
	assert.Nil(t, err)
	assert.NotEqual(t, len(stations), 0)
}

func TestStationFindByRadius(t *testing.T) {
	ur := testHelperStation()
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
		StationName: "station name2",
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
	stations, err := ur.FindInRadius(0, 0, 20, 0, 1000)
	assert.Nil(t, err)
	assert.NotEqual(t, len(stations), 0)
}

func TestStationFindByEmail(t *testing.T) {
	ur := testHelperStation()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    156.12,
		Longitude:   1235.2,
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	s, err := ur.FindByLocation(156.12, 1235.2)
	station.ID = id
	assert.Nil(t, err)
	assert.NotEqual(t, &models.Station{}, s)
	ur.DeleteByID(id, "1234567890123456789012345")
}

func TestStationUpdateByID(t *testing.T) {
	ur := testHelperStation()
	ti := models.GetTimeNow()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    156.12,
		Longitude:   1235.2,
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
			DeleteAt: ti,
		},
	}
	id, err := ur.Create(station)
	assert.Nil(t, err)
	station.Latitude = 156.12
	err = ur.UpdateByID(id, station, "1234567890123456789012345")
	assert.Nil(t, err)
	station.ID = id
	s, err := ur.FindByID(id)
	assert.EqualValues(t, s.Latitude, 156.12)
	//ur.DeleteByID(id)
}

func TestStationDeleteByID(t *testing.T) {
	ur := testHelperStation()
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
	sr := testHelperStation()
	stations, err := sr.Read(0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, stations)
}

func TestStationUpdateRatingByID(t *testing.T) {
	ur := testHelperStation()
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
