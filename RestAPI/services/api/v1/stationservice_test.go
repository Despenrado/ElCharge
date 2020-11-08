package v1

import (
	"log"
	"testing"
	"time"

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
	id, err := s.Station().Create(station)
	if err != nil {
		log.Fatal(err)
	}
	return st, id
}

func TestStationCreate(t *testing.T) {
	st, _ := testHelperStation()
	station := &models.Station{
		Description: "testText",
		StationName: "station name",
		Location:    "156.12 1235.3",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		},
	}
	s, err := st.CreateStation(station)
	assert.Nil(t, err)
	assert.NotEqual(t, st, "")
	st.DeleteByID(s.ID)
}

func TestStationFindByID(t *testing.T) {
	st, id := testHelperStation()
	s, err := st.FindByID(id)
	assert.Nil(t, err)
	assert.NotEqual(t, s, nil)
	s, err = st.FindByID("wrongID")
	assert.NotNil(t, err)
	assert.Nil(t, s)
	st.DeleteByID(id)
}

func TestStationFindByLocation(t *testing.T) {
	st, id := testHelperStation()
	u, err := st.FindByLocation("156.12 1235.2")
	assert.Nil(t, err)
	assert.NotEqual(t, u, nil)
	u, err = st.FindByLocation("wrongLocation")
	assert.NotNil(t, err)
	assert.Nil(t, u)
	st.DeleteByID(id)
}

func TestStationUpdateByID(t *testing.T) {
	st, id := testHelperStation()
	stat, err := st.FindByID(id)
	stat.Location = "Location"
	s, err := st.UpdateByID(id, stat)
	assert.Nil(t, err)
	assert.EqualValues(t, s.Location, "Location")
	assert.EqualValues(t, s.Description, "testText")
	assert.NotEqualValues(t, s.Location, "username_1")
	assert.NotEqualValues(t, s.Description, "2@email.com")
	st.DeleteByID(id)
}

func TestStationDeleteByID(t *testing.T) {
	st, id := testHelperStation()
	stat, err := st.FindByID(id)
	assert.Nil(t, err)
	assert.NotNil(t, stat)
	err = st.DeleteByID(id)
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
