package teststorage

import (
	"errors"
	"time"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StationRepository struct {
	storage *Storage
	db      map[string]*models.Station
}

// NewStationRepository constructor
func NewStationRepository() *StationRepository {
	return &StationRepository{
		db: make(map[string]*models.Station),
	}
}

// Create create and save item in database (map in RAM)
func (r *StationRepository) Create(s *models.Station) (string, error) {
	s.ID = primitive.NewObjectID().Hex()
	r.db[s.ID] = s
	return s.ID, nil
}

// FindByID find item by id and return it (form map in RAM)
func (r *StationRepository) FindByID(id string) (*models.Station, error) {
	s, ok := r.db[id]
	if !ok {
		return nil, utils.ErrRecordNotFound
	}
	return s, nil
}

// FindByEmail find item by email and return it (form map in RAM)
func (r *StationRepository) FindByLocation(location string) (*models.Station, error) {
	for _, s := range r.db {
		if s.Location == location {
			return s, nil
		}
	}
	return nil, utils.ErrRecordNotFound
}

// UpdateByID update item in db (map in RAM)
func (r *StationRepository) UpdateByID(id string, s *models.Station) error {
	s, ok := r.db[id]
	if !ok {
		return utils.ErrRecordNotFound
	}
	if s.UpdateAt != r.db[id].UpdateAt {
		return errors.New("Uptime not equal")
	}
	r.db[id].UpdateAt = time.Now()
	if s.Description != "" {
		r.db[id].Description = s.Description
	}
	if s.Raiting != 0 {
		r.db[id].Raiting = s.Raiting
	}
	if s.StationName != "" {
		r.db[id].StationName = s.StationName
	}
	return nil
}

// DeleteByID delete item by if from db (map in RAM)
func (r *StationRepository) DeleteByID(id string) error {
	delete(r.db, id)
	return nil
}

func (r *StationRepository) Read(skip int, limit int) ([]models.Station, error) {
	stations := make([]models.Station, 0)
	i2 := 0
	for _, v := range r.db {
		i2++
		if i2 >= skip {
			stations = append(stations, *v)
			if i2 < skip+limit {
				break
			}
		}
	}
	return stations, nil
}
