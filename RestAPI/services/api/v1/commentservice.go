package v1

import (
	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/storage"
)

type CommentService struct {
	service *Service
	storage storage.Storage
}

// CommentService connstructor
func NewCommentService(storage storage.Storage) *CommentService {
	return &CommentService{
		storage: storage,
	}
}

func (s *CommentService) CreateComment(sid string, c *models.Comment) (*models.Comment, error) {
	if err := c.Validate(); err != nil {
		return c, err
	}
	if err := c.BeforeCreate(); err != nil {
		return c, err
	}
	id, err := s.storage.Comment().Create(sid, c)
	if err != nil {
		return nil, err
	}
	c, err = s.storage.Comment().FindByID(sid, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// FindByID search station in storage
func (s *CommentService) FindByID(sid string, id string) (*models.Comment, error) {
	c, err := s.storage.Comment().FindByID(sid, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (s *CommentService) FindByUserName(sid string, userName string) (*models.Comment, error) {
	st, err := s.storage.Comment().FindByUserName(sid, userName)
	if err != nil {
		return nil, err
	}
	return st, nil
}

// UpdateByID update storage
func (s *CommentService) UpdateByID(sid string, id string, c *models.Comment) (*models.Comment, error) {
	err := s.storage.Comment().UpdateByID(sid, id, c)
	if err != nil {
		return nil, err
	}
	c, err = s.storage.Comment().FindByID(sid, id)
	if err != nil {
		return nil, err
	}
	return c, nil
}

// DeleteByID delete storage
func (s *CommentService) DeleteByID(sid string, id string) error {
	return s.storage.Comment().DeleteByID(sid, id)
}

func (s *CommentService) Read(sid string, skip int, limit int) ([]models.Comment, error) {
	st, err := s.storage.Station().FindByID(sid)
	if err != nil {
		return nil, err
	}
	comments := make([]models.Comment, 0)
	for i := skip; i < skip+limit && i < len(st.Comments); i++ {
		comments = append(comments, st.Comments[i])
	}
	return comments, nil
}
