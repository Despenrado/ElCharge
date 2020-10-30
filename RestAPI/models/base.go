package models

import "time"

// Model base model of common table structure
type Model struct {
	ID       string    `bson:"_id,omitempty" json:"_id,omitempty"`
	CreateAt time.Time `bson:"create_at,omitempty" json:"create_at,omitempty"`
	UpdateAt time.Time `bson:"updat_at,omitempty" json:"updat_at,omitempty"`
	DeleteAt time.Time `bson:"delete_at,omitempty" json:"delete_at,omitempty"`
}
