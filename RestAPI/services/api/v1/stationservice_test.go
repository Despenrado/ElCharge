package v1

import (
	"log"
	"testing"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	teststorage "github.com/Despenrado/ElCharge/RestAPI/storage/teststore"
	"github.com/stretchr/testify/assert"
)

func TestNewStationService(t *testing.T) {
	ur := teststorage.NewUserRepository()
	sr := teststorage.NewStationRepository()
	cr := teststorage.NewCommentRepository(sr)
	s := teststorage.NewStorage(ur, sr, cr)
	st := NewStationService(s)
	assert.NotNil(t, st)
}

func testHelperStation() (*StationService, string) {
	ur := teststorage.NewUserRepository()
	sr := teststorage.NewStationRepository()
	cr := teststorage.NewCommentRepository(sr)
	s := teststorage.NewStorage(ur, sr, cr)
	st := NewStationService(s)
	ti := models.GetTimeNow()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    15.12,
		Longitude:   12.2,
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	if err := station.BeforeCreate(); err != nil {
		log.Fatal(err)
	}
	id, err := s.Station().Create(station)
	if err != nil {
		log.Fatal(err)
	}
	return st, id
}

func TestStationCreate(t *testing.T) {
	st, _ := testHelperStation()
	ti := models.GetTimeNow()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    16.12,
		Longitude:   12.3,
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	s, err := st.CreateStation(station)
	assert.Nil(t, err)
	assert.NotEqual(t, s, "")
}

func TestStationFindByID(t *testing.T) {
	st, id := testHelperStation()
	s, err := st.FindByID(id)
	assert.Nil(t, err)
	assert.NotEqual(t, s, nil)
	s, err = st.FindByID("wrongID")
	assert.NotNil(t, err)
	assert.Nil(t, s)
}

func TestStationFindByName(t *testing.T) {
	st, _ := testHelperStation()
	u, err := st.FindByName("name")
	assert.Nil(t, err)
	assert.NotEqual(t, u, nil)
}

func TestStationFindByDescription(t *testing.T) {
	st, _ := testHelperStation()
	u, err := st.FindByDescription("text")
	assert.Nil(t, err)
	assert.NotEqual(t, u, nil)
}

func TestStationFindByLocation(t *testing.T) {
	st, _ := testHelperStation()
	u, err := st.FindByLocation(15.12, 12.2)
	assert.Nil(t, err)
	assert.NotEqual(t, u, nil)
	u, err = st.FindByLocation(156.122, 12)
	assert.NotNil(t, err)
	assert.Nil(t, u)
}

func TestStationFindInRadius(t *testing.T) {
	st, _ := testHelperStation()
	ti := models.GetTimeNow()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    0.1,
		Longitude:   0.1,
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	if err := station.BeforeCreate(); err != nil {
		log.Fatal(err)
	}
	_, err := st.CreateStation(station)
	if err != nil {
		log.Fatal(err)
	}
	assert.Nil(t, err)
	ti = models.GetTimeNow()
	station = &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    10,
		Longitude:   14,
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	if err := station.BeforeCreate(); err != nil {
		log.Fatal(err)
	}
	_, err = st.CreateStation(station)
	if err != nil {
		log.Fatal(err)
	}
	assert.Nil(t, err)
	ti = models.GetTimeNow()
	station = &models.Station{
		Description: "testText",
		StationName: "station name",
		OwnerID:     "1234567890123456789012345",
		Latitude:    10,
		Longitude:   0.1,
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	if err := station.BeforeCreate(); err != nil {
		log.Fatal(err)
	}
	_, err = st.CreateStation(station)
	if err != nil {
		log.Fatal(err)
	}
	assert.Nil(t, err)
	u, err := st.FindInRadius(0, 0, 5000, 0, 1000)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(u), 2)
}

func TestStationUpdateByID(t *testing.T) {
	st, id := testHelperStation()
	stat, err := st.FindByID(id)
	stat.Latitude = 153
	s, err := st.UpdateByID(id, stat, "1234567890123456789012345")
	assert.Nil(t, err)
	assert.EqualValues(t, s.Latitude, 153)
	assert.EqualValues(t, s.Description, "testText")
	assert.NotEqualValues(t, s.Latitude, "username_1")
	assert.NotEqualValues(t, s.Description, "2@email.com")
}

func TestStationDeleteByID(t *testing.T) {
	st, id := testHelperStation()
	stat, err := st.FindByID(id)
	assert.Nil(t, err)
	assert.NotNil(t, stat)
	err = st.DeleteByID(id, "1234567890123456789012345")
	assert.Nil(t, err)
	stat, err = st.FindByID(id)
	assert.NotNil(t, err)
	assert.Nil(t, stat)
}

func TestServiceRead(t *testing.T) {
	ur, _ := testHelperStation()
	users, err := ur.Read(0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}

func TestGetLongitude(t *testing.T) {
	assert.Equal(t, getLongitude(181), float64(-179))
	assert.Equal(t, getLongitude(-1), float64(-1))
	assert.Equal(t, getLongitude(0), float64(0))
	assert.Equal(t, getLongitude(18), float64(18))
}

func TestGetLatitude(t *testing.T) {
	assert.Equal(t, getLatitude(91), float64(90))
	assert.Equal(t, getLatitude(-1), float64(-1))
	assert.Equal(t, getLatitude(1), float64(1))
	assert.Equal(t, getLatitude(0), float64(0))
	assert.Equal(t, getLatitude(-92), float64(-90))
}

func TestGetRadius(t *testing.T) {
	assert.GreaterOrEqual(t, getRadius(86), float64(0.76))
	assert.LessOrEqual(t, getRadius(86), float64(0.78))
	assert.GreaterOrEqual(t, getRadius(40078), float64(359))
	assert.LessOrEqual(t, getRadius(40078), float64(360))
	assert.GreaterOrEqual(t, getRadius(40075), float64(359))
	assert.LessOrEqual(t, getRadius(40075), float64(360))
	assert.LessOrEqual(t, getRadius(40070), float64(360))
}
