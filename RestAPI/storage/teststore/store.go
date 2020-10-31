package teststorage

import "github.com/Despenrado/ElCharge/RestAPI/storage"

// Storage ...
type Storage struct {
	userRepository *UserRepository
}

// User ...
func (s *Storage) User() storage.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		storage: s,
	}

	return s.userRepository
}
