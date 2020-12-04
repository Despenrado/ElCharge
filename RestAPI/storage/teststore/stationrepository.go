package teststorage

import (
	"errors"
	"log"
	"regexp"

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
func (r *StationRepository) FindByLocation(positionX float64, positionY float64) (*models.Station, error) {
	for _, s := range r.db {
		if s.Latitude == positionX && s.Longitude == positionY {
			return s, nil
		}
	}
	return nil, utils.ErrRecordNotFound
}

func (r *StationRepository) FindByName(name string) ([]models.Station, error) {
	regex, _ := regexp.Compile(".*" + name + ".*")
	st := []models.Station{}
	for _, s := range r.db {
		if regex.MatchString(s.StationName) {
			st = append(st, *s)
		}
	}
	return st, nil
}

func (r *StationRepository) FindByDescription(text string) ([]models.Station, error) {
	regex, _ := regexp.Compile(".*" + text + ".*")
	st := []models.Station{}
	for _, s := range r.db {
		if regex.MatchString(s.Description) {
			st = append(st, *s)
		}
	}
	return st, nil
}

func (r *StationRepository) FindInRadius(latitude float64, longitude float64, radius float64, skip int, limit int) ([]models.Station, error) {
	var stations []models.Station
	i2 := 0
	for _, v := range r.db {
		log.Println(i2)
		if latitude-radius <= v.Latitude && latitude+radius >= v.Latitude && longitude-radius <= v.Longitude && longitude+radius >= v.Longitude {
			i2++
			if i2 >= skip {
				stations = append(stations, *v)
				if i2 >= skip+limit {
					break
				}
			}
		}
	}
	return stations, nil
}

// UpdateByID update item in db (map in RAM)
func (r *StationRepository) UpdateByID(id string, s *models.Station, ownerID string) error {
	if r.db[id].OwnerID != ownerID {
		return errors.New("userID and ownerID not equel")
	}
	s, ok := r.db[id]
	if !ok {
		return utils.ErrRecordNotFound
	}
	if s.UpdateAt != r.db[id].UpdateAt {
		return errors.New("Uptime not equal")
	}
	r.db[id].UpdateAt = models.GetTimeNow()
	if s.Description != "" {
		r.db[id].Description = s.Description
	}
	if s.Rating != 0 {
		r.db[id].Rating = s.Rating
	}
	if s.StationName != "" {
		r.db[id].StationName = s.StationName
	}
	return nil
}

// DeleteByID delete item by if from db (map in RAM)
func (r *StationRepository) DeleteByID(id string, ownerID string) error {
	if r.db[id].OwnerID == ownerID {
		delete(r.db, id)
		return nil
	}
	return errors.New("userID and ownerID not equel")
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

func (r *StationRepository) UpdateRaitingByID(id string) error {
	comments := r.db[id].Comments
	var sum float32
	sum = 0
	for _, v := range comments {
		sum += v.Rating
	}
	r.db[id].Rating = sum / float32(len(comments))
	return nil
}
