package mongostorage

import (
	"log"
	"testing"
	"time"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/stretchr/testify/assert"
)

func testHelperUser() *UserRepository {
	db, err := ConnectToDB("mongodb://127.0.0.1:27017", "elCharge")
	if err != nil {
		log.Fatal(err)
	}
	ur := NewUserRepository(ConfigureRepository(db, "user"))
	return ur
}

func TestCreate(t *testing.T) {
	ur := testHelperUser()
	user := &models.User{
		UserName: "username_1",
		Email:    "1@email.com",
		Password: "passwoed_1",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
		},
	}
	id, err := ur.Create(user)
	assert.Nil(t, err)
	assert.NotEqual(t, id, "")
	ur.DeleteByID(id)
}

func TestFindByID(t *testing.T) {
	ur := testHelperUser()
	user := &models.User{
		UserName: "username_1",
		Email:    "1@email.com",
		Password: "passwoed_1",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
			DeleteAt: time.Now(),
		},
	}
	id, err := ur.Create(user)
	assert.Nil(t, err)
	u, err := ur.FindByID(id)
	assert.NotNil(t, u)
	ur.DeleteByID(id)
}

func TestFindByEmail(t *testing.T) {
	ur := testHelperUser()
	user := &models.User{
		UserName: "username_1",
		Email:    "1@email.com",
		Password: "passwoed_1",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
			DeleteAt: time.Now(),
		},
	}
	id, err := ur.Create(user)
	assert.Nil(t, err)
	u, err := ur.FindByEmail("1@email.com")
	assert.NotNil(t, u)
	ur.DeleteByID(id)
}

func TestUpdateByID(t *testing.T) {
	ur := testHelperUser()
	user := &models.User{
		UserName: "username_1",
		Email:    "1@email.com",
		Password: "passwoed_1",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
			DeleteAt: time.Now(),
		},
	}
	id, err := ur.Create(user)
	assert.Nil(t, err)
	user.Email = "2@email.com"
	err = ur.UpdateByID(id, user)
	assert.Nil(t, err)
	u, err := ur.FindByID(id)
	assert.NotNil(t, u)
	assert.Equal(t, u.Email, user.Email)
	ur.DeleteByID(id)
}

func TestDeleteByID(t *testing.T) {
	ur := testHelperUser()
	user := &models.User{
		UserName: "username_1",
		Email:    "1@email.com",
		Password: "passwoed_1",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
			DeleteAt: time.Now(),
		},
	}
	id, err := ur.Create(user)
	assert.Nil(t, err)
	err = ur.DeleteByID(id)
	assert.Nil(t, err)
	u, err := ur.FindByID(id)
	assert.Nil(t, u)
}

func TestUserRead(t *testing.T) {
	ur := testHelperUser()
	users, err := ur.Read(0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}
