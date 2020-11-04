package storage

import "github.com/Despenrado/ElCharge/RestAPI/models"

// UserRepository ...
type UserRepository interface {
	Create(*models.User) (string, error)
	FindByID(string) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	UpdateByID(string, *models.User) (*models.User, error)
	DeleteByID(string) error
}
