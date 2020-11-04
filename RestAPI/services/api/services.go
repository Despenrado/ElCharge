package api

import "github.com/Despenrado/ElCharge/RestAPI/models"

type UserService interface {
	CreateUser(*models.User) (*models.User, error)
	Login(*models.User) (*models.User, error)
	FindByID(string) (*models.User, error)
	UpdateByID(string, *models.User) (*models.User, error)
	DeleteByID(id string) error
}
