package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Station struct {
	ID          string    `bson:"_id,omitempty" json:"_id,omitempty"`
	StationName string    `bson:"station_name,omitempty" json:"station_name,omitempty"`
	Raiting     int8      `bson:"raiting,omitempty" json:"raiting,omitempty"`
	Description string    `bson:"description,omitempty" json:"description,omitempty"`
	Comments    []Comment `bson:"comments,omitempty" json:"comments,omitempty"`
	Location    string    `bson:"location,omitempty" json:"location,omitempty"`
	Model
}

// Validate ...
func (s *Station) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.StationName, validation.Required, validation.Length(2, 100)),
		validation.Field(&s.Description, validation.Required, validation.Length(5, 512)),
		validation.Field(&s.Location, validation.Required, validation.Length(8, 256)),
	)
}

// BeforeCreate some manipulation whith item before seve to db
func (s *Station) BeforeCreate() error {
	s.Model.BeforeCreate()
	return nil
}
