package mongostorage

import (
	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/storage"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository ...
type UserRepository struct {
	storage *Storage
	col     mongo.Collection
}

// Create ...
func (r *UserRepository) Create(u *models.User) error {

	return nil
}

// Find ...
func (r *UserRepository) Find(id string) (*models.User, error) {

	return nil, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {

	return nil, storage.ErrRecordNotFound
}
