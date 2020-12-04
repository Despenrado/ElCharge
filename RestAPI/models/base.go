package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"golang.org/x/crypto/bcrypt"
)

// Model base model of common table structure
type Model struct {
	CreateAt time.Time `bson:"create_at,omitempty" json:"create_at,omitempty"`
	UpdateAt time.Time `bson:"update_at,omitempty" json:"update_at,omitempty"`
	DeleteAt time.Time `bson:"delete_at,omitempty" json:"delete_at,omitempty"`
}

func isRequired(b bool) validation.RuleFunc {
	return func(val interface{}) error {
		if b {
			return validation.Validate(val, validation.Required)
		}
		return nil
	}
}

// EncryptString ...
func EncryptString(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// BeforeCreate some manipulation whith item before seve to db
func (m *Model) BeforeCreate() error {
	t := GetTimeNow()
	m.CreateAt = t
	m.UpdateAt = t
	return nil
}

func GetTimeNow() time.Time {
	t := time.Now()
	return time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), 0, t.Location())
}
