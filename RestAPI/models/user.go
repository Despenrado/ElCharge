package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User structure
type User struct {
	ID       string `bson:"_id,omitempty" json:"_id,omitempty"`
	UserName string `bson:"user_name,omitempty" json:"user_name,omitempty"`
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
	Model
}

// Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.UserName, validation.By(isRequired(u.UserName != "")), validation.Length(2, 100)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(isRequired(u.Password == "")), validation.Length(8, 50)),
	)
}

// BeforeCreate some manipulation whith item before seve to db
func (u *User) BeforeCreate() error {
	if len(u.Password) <= 0 {
		return nil
	}
	enc, err := EncryptString(u.Password)
	if err != nil {
		return err
	}
	u.Password = enc
	u.Model.BeforeCreate()
	return nil
}

// VerifyPassword is pasword equels to db password
func (u *User) VerifyPassword(p string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p)) == nil
}

// Sanitize delete some fields (make dto)
func (u *User) Sanitize() {
	u.Password = ""
	u.CreateAt = time.Time{}
}
