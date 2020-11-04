package v1

import (
	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/storage"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
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
	_, err := s.storage.User().FindByEmail(u.Email)
	if err != utils.ErrRecordNotFound {
		if err != nil {
			return nil, err
		}
		return u, utils.ErrRecordAlreadyExists
	}
	if err := u.BeforeCreate(); err != nil {
		return u, err
	}
	id, err := s.storage.User().Create(u)
	if err != nil {
		return nil, err
	}
	u, err = s.storage.User().FindByID(id)
	if err != nil {
		return nil, err
	}
	u.Sanitize()
	return u, nil
}

func (s *UserService) Login(u *models.User) (*models.User, error) {
	if err := u.Validate(); err != nil {
		return u, err
	}
	u2, err := s.storage.User().FindByEmail(u.Email)
	if err != nil {
		return u, err
	}
	if !u2.VerifyPassword(u.Password) {
		return u, utils.ErrIncorrectEmailOrPassword
	}
	u2.Sanitize()
	return u2, nil
}

func (s *UserService) FindByID(id string) (*models.User, error) {
	u, err := s.storage.User().FindByID(id)
	if err != nil {
		return nil, err
	}
	u.Sanitize()
	return u, nil
}

func (s *UserService) UpdateByID(id string, u *models.User) (*models.User, error) {
	u, err := s.storage.User().UpdateByID(id, u)
	if err != nil {
		return nil, err
	}
	u, err = s.storage.User().FindByID(id)
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (s *UserService) DeleteByID(id string) error {
	return s.storage.User().DeleteByID(id)
}
