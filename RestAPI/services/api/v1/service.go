package v1

import "github.com/Despenrado/ElCharge/RestAPI/services/api"

type Service struct {
	userService *UserService
}

func NewService(us *UserService) *Service {
	s := &Service{
		userService: us,
	}
	s.userService.service = s
	return s
}

// User ...
func (s *Service) User() api.UserService {
	if s.userService != nil {
		return s.userService
	}
	s.userService = &UserService{
		service: s,
	}

	return s.userService
}
