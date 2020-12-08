package v1

import (
	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/storage"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
)

// UserService struct
type UserService struct {
	service *Service
	storage storage.Storage
}

// NewUserService connstructor
func NewUserService(storage storage.Storage) *UserService {
	return &UserService{
		storage: storage,
	}
}

// CreateUser save user to storage
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

// Login find user that have email and password like 'u'
func (s *UserService) Login(u *models.User) (*models.User, error) {
	if err := u.Validate(); err != nil {
		return nil, err
	}
	u2, err := s.storage.User().FindByEmail(u.Email)
	if err != nil {
		return nil, err
	}
	if !u2.VerifyPassword(u.Password) {
		return nil, utils.ErrIncorrectEmailOrPassword
	}
	u2.Sanitize()
	return u2, nil
}

// FindByID search user in storage
func (s *UserService) FindByID(id string) (*models.User, error) {
	u, err := s.storage.User().FindByID(id)
	if err != nil {
		return nil, err
	}
	u.Sanitize()
	return u, nil
}

// UpdateByID update user
func (s *UserService) UpdateByID(id string, u *models.User) (*models.User, error) {
	if u.Password != "" {
		tmp, err := models.EncryptString(u.Password)
		u.Password = tmp
		if err != nil {
			return nil, err
		}
	}
	err := s.storage.User().UpdateByID(id, u)
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

// DeleteByID delete user
func (s *UserService) DeleteByID(id string) error {
	return s.storage.User().DeleteByID(id)
}

func (s *UserService) Read(skip int, limit int) ([]models.User, error) {
	us, err := s.storage.User().Read(skip, limit)
	if err != nil {
		return nil, err
	}
	return us, nil
}
