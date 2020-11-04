package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	testCase := []struct {
		name    string
		user    *User
		isValid bool
	}{
		{
			name: "simple",
			user: &User{
				UserName: "valid",
				Email:    "valid@email.com",
				Password: "validpass",
			},
			isValid: true,
		},
		{
			name: "without username",
			user: &User{
				Email:    "valid@email.com",
				Password: "validpass",
			},
			isValid: true,
		},
		{
			name: "without email",
			user: &User{
				UserName: "valid",
				Password: "validpass",
			},
			isValid: false,
		},
		{
			name: "without password",
			user: &User{
				UserName: "valid",
				Email:    "valid@email.com",
			},
			isValid: false,
		},
		{
			name: "wrong email",
			user: &User{
				UserName: "valid",
				Email:    "wrong",
				Password: "validpass",
			},
			isValid: false,
		},
		{
			name: "wrong password",
			user: &User{
				UserName: "valid",
				Email:    "valid@email.com",
				Password: "wrong",
			},
			isValid: false,
		},
		{
			name: "wrong username",
			user: &User{
				UserName: "n",
				Email:    "valid@email.com",
				Password: "validpass",
			},
			isValid: false,
		},
	}

	for _, item := range testCase {
		t.Run(item.name, func(t *testing.T) {
			if item.isValid {
				assert.NoError(t, item.user.Validate())
			} else {
				assert.Error(t, item.user.Validate())
			}
		})
	}
}

func TestBeforeCreate(t *testing.T) {
	user := &User{
		UserName: "valid",
		Email:    "valid@email.com",
		Password: "validpass",
	}
	assert.NoError(t, user.BeforeCreate())
	assert.NotEqualValues(t, user.Password, "validpass")
	assert.NotNil(t, user.UpdateAt)
	assert.NotNil(t, user.CreateAt)
}

func TestSanitize(t *testing.T) {
	user := &User{
		UserName: "valid",
		Email:    "valid@email.com",
		Password: "validpass",
	}
	assert.NoError(t, user.BeforeCreate())
	user.Sanitize()
	assert.EqualValues(t, user.Password, "")
	assert.NotEqualValues(t, user.Email, "")
	assert.NotEqualValues(t, user.UserName, "")
	assert.Equal(t, user.CreateAt, time.Time{})
	assert.NotEqual(t, user.UpdateAt, time.Time{})
}

func TestVerifyPassword(t *testing.T) {
	user := &User{
		UserName: "valid",
		Email:    "valid@email.com",
		Password: "validpass",
	}
	assert.NoError(t, user.BeforeCreate())
	assert.EqualValues(t, user.VerifyPassword("validpass"), true)
	assert.NotEqualValues(t, user.VerifyPassword("notvalidpass"), true)
}
