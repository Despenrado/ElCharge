package teststorage

import (
	"time"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepository struct for testing
type UserRepository struct {
	storage *Storage
	db      map[string]*models.User
}

// NewUserRepository constructor
func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: make(map[string]*models.User),
	}
}

// Create create and save item in database (map in RAM)
func (r *UserRepository) Create(u *models.User) (string, error) {
	u.ID = primitive.NewObjectID().Hex()
	r.db[u.ID] = u
	return u.ID, nil
}

// FindByID find item by id and return it (form map in RAM)
func (r *UserRepository) FindByID(id string) (*models.User, error) {
	u, ok := r.db[id]
	if !ok {
		return nil, utils.ErrRecordNotFound
	}
	return u, nil
}

// FindByEmail find item by email and return it (form map in RAM)
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	for _, u := range r.db {
		if u.Email == email {
			return u, nil
		}
	}
	return nil, utils.ErrRecordNotFound
}

// UpdateByID update item in db (map in RAM)
func (r *UserRepository) UpdateByID(id string, u *models.User) (*models.User, error) {
	u, ok := r.db[id]
	if !ok {
		return nil, utils.ErrRecordNotFound
	}
	r.db[id].UpdateAt = time.Now()
	if u.Email != "" {
		r.db[id].Email = u.Email
	}
	if u.UserName != "" {
		r.db[id].UserName = u.UserName
	}
	if u.Password != "" {
		var err error
		r.db[id].Password, err = models.EncryptString(u.Password)
		if err != nil {
			return nil, err
		}
	}
	return r.db[id], nil
}

// DeleteByID delete item by if from db (map in RAM)
func (r *UserRepository) DeleteByID(id string) error {
	delete(r.db, id)
	return nil
}
