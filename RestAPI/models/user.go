package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

// User structure
type User struct {
	Model
	UserName string `bson:"user_name,omitempty" json:"user_name,omitempty"`
	Email    string `bson:"email,omitempty" json:"email,omitempty"`
	Password string `bson:"password,omitempty" json:"password,omitempty"`
}

// Validate ...
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.UserName, validation.Required, validation.Length(2, 100)),
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Password, validation.By(isRequired(u.Password == "")), validation.Length(8, 50)),
	)
}

func (u *User) BeforeCreate() error {
	if len(u.Password) <= 0 {
		return nil
	}
	enc, err := encryptString(u.Password)
	if err != nil {
		return err
	}
	u.Password = enc
	t := time.Now()
	u.CreateAt = t
	u.UpdateAt = t
	return nil
}

func (u *User) VerifyPassword(p string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p)) == nil
}

func (u *User) Sanitize() {
	u.Password = ""
	u.CreateAt = time.Time{}
}

func encryptString(str string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
