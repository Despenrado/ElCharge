package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

// Model base model of common table structure
type Model struct {
	ID       string    `bson:"_id,omitempty" json:"_id,omitempty"`
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
