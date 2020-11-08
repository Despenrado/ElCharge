package v1

import "github.com/Despenrado/ElCharge/RestAPI/services/api"

// Service service struct
type Service struct {
	userService    *UserService
	stationService *StationService
	commentService *CommentService
}

// NewService constructor
func NewService(us *UserService, ss *StationService, cs *CommentService) *Service {
	s := &Service{
		userService:    us,
		stationService: ss,
		commentService: cs,
	}
	s.userService.service = s
	return s
}

// User return UserService
func (s *Service) User() api.UserService {
	if s.userService != nil {
		return s.userService
	}
	s.userService = &UserService{
		service: s,
	}
	return s.userService
}

// Station return UserService
func (s *Service) Station() api.StationService {
	if s.stationService != nil {
		return s.stationService
	}
	s.stationService = &StationService{
		service: s,
	}
	return s.stationService
}

// Comment return UserService
func (s *Service) Comment() api.CommentService {
	if s.commentService != nil {
		return s.commentService
	}
	s.commentService = &CommentService{
		service: s,
	}
	return s.commentService
}
