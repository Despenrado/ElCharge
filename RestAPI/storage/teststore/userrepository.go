package teststorage

import (
	"github.com/Despenrado/ElCharge/models"
	"github.com/Despenrado/ElCharge/storage"
)

// UserRepository ...
type UserRepository struct {
	storage *Storage
	db      map[string]*models.User
}

// Create ...
func (r *UserRepository) Create(u *models.User) error {
	// if err := u.Validate(); err != nil {
	// 	return err
	// }

	// if err := u.BeforeCreate(); err != nil {
	// 	return err
	// }

	// u.ID = len(r.users) + 1
	// r.users[u.ID] = u

	// return nil
	return nil
}

// Find ...
func (r *UserRepository) Find(id int) (*models.User, error) {
	// u, ok := r.users[id]
	// if !ok {
	// 	return nil, store.ErrRecordNotFound
	// }

	// return u, nil
	return nil, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	// for _, u := range r.users {
	// 	if u.Email == email {
	// 		return u, nil
	// 	}
	// }

	return nil, storage.ErrRecordNotFound
}
