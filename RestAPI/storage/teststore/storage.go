package teststorage

import "github.com/Despenrado/ElCharge/RestAPI/storage"

// Storage ...
type Storage struct {
	userRepository    *UserRepository
	stationRepository *StationRepository
	commentRepository *CommentRepository
}

// NewStorage constructor
func NewStorage(ur *UserRepository, sr *StationRepository, cr *CommentRepository) *Storage {
	s := &Storage{
		userRepository:    ur,
		stationRepository: sr,
		commentRepository: cr,
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

// Station ...
func (s *Storage) Station() storage.StationRepository {
	if s.stationRepository != nil {
		return s.stationRepository
	}
	s.stationRepository = &StationRepository{
		storage: s,
	}
	return s.stationRepository
}

// Comment ...
func (s *Storage) Comment() storage.CommentRepository {
	if s.commentRepository != nil {
		return s.commentRepository
	}
	s.commentRepository = &CommentRepository{
		storage: s,
	}
	return s.commentRepository
}
