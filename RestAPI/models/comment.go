package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
	ID       string  `bson:"_id,omitempty" json:"_id,omitempty"`
	UserID   string  `bson:"user_id,omitempty" json:"user_id,omitempty"`
	UserName string  `bson:"user_name,omitempty" json:"user_name,omitempty"`
	Text     string  `bson:"text,omitempty" json:"text,omitempty"`
	Rating   float32 `bson:"rating,omitempty,truncate" json:"rating,omitempty,truncate"`
	Model
}

// Validate ...
func (c *Comment) Validate() error {
	return validation.ValidateStruct(
		c,
		validation.Field(&c.UserName, validation.Required, validation.Length(2, 100)),
		validation.Field(&c.UserID, validation.Required, validation.Length(20, 30)),
		validation.Field(&c.Text, validation.Required, validation.Length(8, 256)),
	)
}

// BeforeCreate some manipulation whith item before seve to db
func (c *Comment) BeforeCreate() error {
	c.ID = primitive.NewObjectID().Hex()
	c.Model.BeforeCreate()
	return nil
}
