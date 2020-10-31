package mongostorage

import "github.com/Despenrado/ElCharge/RestAPI/storage"

// Storage ...
type Storage struct {
	userRepository *UserRepository
}

func NewStorage(ur *UserRepository) *Storage {
	s := &Storage{
		userRepository: ur,
	}
	s.userRepository.storage = s
	return s
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
