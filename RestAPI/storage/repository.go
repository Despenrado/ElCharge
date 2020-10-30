package storage

import "github.com/Despenrado/ElCharge/RestAPI/models"

// UserRepository ...
type UserRepository interface {
	Create(*models.User) error
	Find(string) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	Delete(string) error
}
