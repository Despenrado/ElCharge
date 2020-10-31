package api

import "github.com/Despenrado/ElCharge/RestAPI/models"

type UserService interface {
	CreateUser(*models.User) (*models.User, error)
}
