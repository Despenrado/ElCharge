package teststorage

import (
	"testing"
	"time"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ur := NewUserRepository()
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
	ur := NewUserRepository()
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
	ur := NewUserRepository()
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
	ur := NewUserRepository()
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
	assert.Nil(t, err)
	assert.Equal(t, u, user)
	assert.EqualValues(t, u.Email, "2@email.com")
}

func TestDeleteByID(t *testing.T) {
	ur := NewUserRepository()
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
	ur := NewUserRepository()
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
	user = &models.User{
		UserName: "username_1",
		Email:    "1@email.com",
		Password: "passwoed_1",
		Model: models.Model{
			UpdateAt: time.Now(),
			CreateAt: time.Now(),
			DeleteAt: time.Now(),
		},
	}
	_, err = ur.Create(user)
	assert.Nil(t, err)
	users, err := ur.Read(0, 10)
	assert.Nil(t, err)
	assert.NotNil(t, users)
}
