package v1

import (
	"log"
	"testing"
	"time"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	teststorage "github.com/Despenrado/ElCharge/RestAPI/storage/teststore"
	"github.com/stretchr/testify/assert"
)

func TestNewUserService(t *testing.T) {
	ur := teststorage.NewUserRepository()
	st := teststorage.NewStorage(ur)
	us := NewUserService(st)
	assert.NotNil(t, us)
}

func testHelper() (*UserService, string) {
	ur := teststorage.NewUserRepository()
	st := teststorage.NewStorage(ur)
	us := NewUserService(st)
	user := &models.User{
		UserName: "username_1",
		Email:    "1@email.com",
		Password: "passwoed_1",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
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
	user := &models.User{
		UserName: "username_2",
		Email:    "2@email.com",
		Password: "passwoed_2",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
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
	user.Password = "passwoed_2"
	u, err := us.UpdateByID(id, user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.UserName, "username_2")
	assert.EqualValues(t, u.Email, "1@email.com")
	assert.EqualValues(t, u.VerifyPassword("passwoed_2"), true)
	assert.NotEqualValues(t, u.UserName, "username_1")
	assert.NotEqualValues(t, u.Email, "2@email.com")
	assert.NotEqualValues(t, u.VerifyPassword("passwoed_1"), true)
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
