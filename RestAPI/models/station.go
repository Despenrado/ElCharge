package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type Station struct {
	ID          string    `bson:"_id,omitempty" json:"_id,omitempty"`
	StationName string    `bson:"station_name,omitempty" json:"station_name,omitempty"`
	OwnerID     string    `bson:"owner_id,omitempty" json:"owner_id,omitempty"`
	Rating      float32   `bson:"rating,truncate" json:"rating,truncate"`
	Description string    `bson:"description,omitempty" json:"description,omitempty"`
	Comments    []Comment `bson:"comments,omitempty" json:"comments,omitempty"`
	Latitude    float64   `bson:"latitude" json:"latitude"`
	Longitude   float64   `bson:"longitude" json:"longitude"`
	Model
}

// Validate ...
func (s *Station) Validate() error {
	return validation.ValidateStruct(
		s,
		validation.Field(&s.StationName, validation.Required, validation.Length(2, 100)),
		validation.Field(&s.Description, validation.Required, validation.Length(5, 512)),
		validation.Field(&s.OwnerID, validation.Required, validation.Length(20, 30)),
		validation.Field(&s.Latitude, validation.Required, validation.Max(float64(90))),
		validation.Field(&s.Latitude, validation.Required, validation.Min(float64(-90))),
		validation.Field(&s.Longitude, validation.Required, validation.Max(float64(180))),
		validation.Field(&s.Longitude, validation.Required, validation.Min(float64(-180))),
	)
}

// BeforeCreate some manipulation whith item before seve to db
func (s *Station) BeforeCreate() error {
	s.Model.BeforeCreate()
	s.Rating = 0
	return nil
}
