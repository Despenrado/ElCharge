package v1

import (
	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/storage"
)

type UserService struct {
	service *Service
	storage storage.Storage
}

func NewUserService(storage storage.Storage) *UserService {
	return &UserService{
		storage: storage,
	}
}

func (s *UserService) CreateUser(u *models.User) (*models.User, error) {
	if err := u.Validate(); err != nil {
		return u, err
	}
	if err := u.BeforeCreate(); err != nil {
		return u, err
	}
	u.Sanitize()
	return u, nil
}
