package teststorage

import (
	"errors"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentRepository struct {
	storage           *Storage
	stationRepository *StationRepository
}

// NewCommentRepository constructor
func NewCommentRepository(sr *StationRepository) *CommentRepository {
	return &CommentRepository{
		stationRepository: sr,
	}
}

// Create create and save item in database (map in RAM)
func (r *CommentRepository) Create(sid string, c *models.Comment) (string, error) {
	c.ID = primitive.NewObjectID().Hex()
	r.stationRepository.db[sid].Comments = append([]models.Comment{*c}, r.stationRepository.db[sid].Comments...)
	return c.ID, nil
}

// FindByID find item by id and return it (form map in RAM)
func (r *CommentRepository) FindByID(sid string, id string) (*models.Comment, error) {
	for _, c := range r.stationRepository.db[sid].Comments {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, utils.ErrRecordNotFound
}

// FindByUserName find item by user name and return it (form map in RAM)
func (r *CommentRepository) FindByUserName(sid string, uname string) (*models.Comment, error) {
	for _, c := range r.stationRepository.db[sid].Comments {
		if c.UserName == uname {
			return &c, nil
		}
	}
	return nil, utils.ErrRecordNotFound
}

// UpdateByID update item in db (map in RAM)
func (r *CommentRepository) UpdateByID(sid string, id string, c *models.Comment) error {
	for i, v := range r.stationRepository.db[sid].Comments {
		if v.ID == id {
			if r.stationRepository.db[sid].Comments[i].UpdateAt != c.UpdateAt {
				return errors.New("Uptime not equal")
			}
			r.stationRepository.db[sid].Comments[i].UpdateAt = models.GetTimeNow()
			if c.Text != "" {
				r.stationRepository.db[sid].Comments[i].Text = c.Text
			}
			if c.Rating != 0 {
				r.stationRepository.db[sid].Comments[i].Rating = c.Rating
				r.stationRepository.db[sid].Comments[i].Rating = c.Rating
			}
			return nil
		}
	}
	return utils.ErrRecordNotFound
}

// DeleteByID delete item by if from db (map in RAM)
func (r *CommentRepository) DeleteByID(sid string, id string) error {
	for i, c := range r.stationRepository.db[sid].Comments {
		if c.ID == id {
			r.stationRepository.db[sid].Comments = append(r.stationRepository.db[sid].Comments[:i], r.stationRepository.db[sid].Comments[i+1:]...)
			return nil
		}
	}
	return utils.ErrRecordNotFound
}
