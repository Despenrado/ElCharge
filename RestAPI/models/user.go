package models

// User structure
type User struct {
	Model
	UserName string `bson:"user_name,omitempty" json:"user_name,omitempty"`
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}
