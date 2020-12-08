package v1

import (
	"log"
	"testing"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	teststorage "github.com/Despenrado/ElCharge/RestAPI/storage/teststore"
	"github.com/stretchr/testify/assert"
)

func TestNewUserService(t *testing.T) {
	ur := teststorage.NewUserRepository()
	sr := teststorage.NewStationRepository()
	cr := teststorage.NewCommentRepository(sr)
	st := teststorage.NewStorage(ur, sr, cr)
	us := NewUserService(st)
	assert.NotNil(t, us)
}

func testHelper() (*UserService, string) {
	ur := teststorage.NewUserRepository()
	sr := teststorage.NewStationRepository()
	cr := teststorage.NewCommentRepository(sr)
	st := teststorage.NewStorage(ur, sr, cr)
	us := NewUserService(st)
	ti := models.GetTimeNow()
	user := &models.User{
		UserName: "username_1",
		Email:    "1@email.com",
		Password: "passwoed_1",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	if err := user.BeforeCreate(); err != nil {
		log.Fatal(err)
	}
	id, err := st.User().Create(user)
	if err != nil {
		log.Fatal(err)
	}
	return us, id
}

func TestUserCreate(t *testing.T) {
	us, _ := testHelper()
	ti := models.GetTimeNow()
	user := &models.User{
		UserName: "username_2",
		Email:    "2@email.com",
		Password: "passwoed_2",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
		},
	}
	id, err := us.CreateUser(user)
	assert.Nil(t, err)
	assert.NotEqual(t, id, "")
}

func TestLogin(t *testing.T) {
	us, _ := testHelper()
	user := &models.User{
		Email:    "1@email.com",
		Password: "passwoed_1",
	}
	u, err := us.Login(user)
	assert.Nil(t, err)
	assert.NotEqual(t, u, nil)
	user.Password = "wrongPass"
	u, err = us.Login(user)
	assert.NotNil(t, err)
	assert.Nil(t, u)
}

func TestFindByID(t *testing.T) {
	us, id := testHelper()
	u, err := us.FindByID(id)
	assert.Nil(t, err)
	assert.NotEqual(t, u, nil)
	u, err = us.FindByID("wrongID")
	assert.NotNil(t, err)
	assert.Nil(t, u)
}

func TestUpdateByID(t *testing.T) {
	us, id := testHelper()
	user, err := us.FindByID(id)
	user.UserName = "username_2"
	u, err := us.UpdateByID(id, user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.UserName, "username_2")
	assert.EqualValues(t, u.Email, "1@email.com")
	assert.NotEqualValues(t, u.UserName, "username_1")
	assert.NotEqualValues(t, u.Email, "2@email.com")
}

func TestDeleteByID(t *testing.T) {
	us, id := testHelper()
	user, err := us.FindByID(id)
	assert.Nil(t, err)
	assert.NotNil(t, user)
	err = us.DeleteByID(id)
	assert.Nil(t, err)
	user, err = us.FindByID(id)
	assert.NotNil(t, err)
	assert.Nil(t, user)
}
func TestUserRead(t *testing.T) {
	ur, _ := testHelper()
	ti := models.GetTimeNow()
	user := &models.User{
		UserName: "username_2",
		Email:    "2@email.com",
		Password: "passwoed_1",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
			DeleteAt: ti,
		},
	}
	_, err := ur.CreateUser(user)
	assert.Nil(t, err)
	ti = models.GetTimeNow()
	user = &models.User{
		UserName: "username_3",
		Email:    "3@email.com",
		Password: "passwoed_1",
		Model: models.Model{
			UpdateAt: ti,
			CreateAt: ti,
			DeleteAt: ti,
		},
	}
	_, err = ur.CreateUser(user)
	assert.Nil(t, err)
	users, err := ur.Read(0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}
