package teststorage

import (
	"strconv"
	"testing"
	"time"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/stretchr/testify/assert"
)

func testHelper() *UserRepository {
	ur := NewUserRepository()
	for i := 0; i < 5; i++ {
		ur.db[strconv.Itoa(i)] = &models.User{
			UserName: strconv.Itoa(i),
			Email:    strconv.Itoa(i),
			Password: strconv.Itoa(i),
			Model: models.Model{
				UpdateAt: time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC),
				CreateAt: time.Now(),
			},
		}
	}
	return ur
}

func TestCreate(t *testing.T) {
	ur := testHelper()
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
	user.ID = id
	assert.Equal(t, ur.db[id], user)
}

func TestFindByID(t *testing.T) {
	ur := testHelper()
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
	assert.Equal(t, u, user)
}

func TestFindByEmail(t *testing.T) {
	ur := testHelper()
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
	_, err := ur.Create(user)
	assert.Nil(t, err)
	u, err := ur.FindByEmail("1@email.com")
	assert.Equal(t, u, user)
}

func TestUpdateByID(t *testing.T) {
	ur := testHelper()
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
	u, err := ur.UpdateByID(id, user)
	assert.Nil(t, err)
	assert.EqualValues(t, u.Email, "2@email.com")
}

func TestDeleteByID(t *testing.T) {
	ur := testHelper()
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
