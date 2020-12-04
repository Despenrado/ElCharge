package v1

import (
	"math"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/storage"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
)

type StationService struct {
	service *Service
	storage storage.Storage
}

// NewStationService connstructor
func NewStationService(storage storage.Storage) *StationService {
	return &StationService{
		storage: storage,
	}
}

func (s *StationService) CreateStation(st *models.Station) (*models.Station, error) {
	if err := st.Validate(); err != nil {
		return st, err
	}
	_, err := s.storage.Station().FindByLocation(st.Latitude, st.Longitude)
	if err != utils.ErrRecordNotFound {
		if err != nil {
			return nil, err
		}
		return st, utils.ErrRecordAlreadyExists
	}
	if err := st.BeforeCreate(); err != nil {
		return st, err
	}
	id, err := s.storage.Station().Create(st)
	if err != nil {
		return nil, err
	}
	st, err = s.storage.Station().FindByID(id)
	if err != nil {
		return nil, err
	}
	return st, nil
}

// FindByID search station in storage
func (s *StationService) FindByID(id string) (*models.Station, error) {
	st, err := s.storage.Station().FindByID(id)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (s *StationService) FindByName(name string) ([]models.Station, error) {
	st, err := s.storage.Station().FindByName(name)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (s *StationService) FindByDescription(text string) ([]models.Station, error) {
	st, err := s.storage.Station().FindByName(text)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (s *StationService) FindByLocation(latitude float64, longitude float64) (*models.Station, error) {
	st, err := s.storage.Station().FindByLocation(getLatitude(latitude), getLongitude(longitude))
	if err != nil {
		return nil, err
	}
	return st, nil
}

func (s *StationService) FindInRadius(latitude float64, longitude float64, radius int, skip int, limit int) ([]models.Station, error) {
	st, err := s.storage.Station().FindInRadius(getLatitude(latitude), getLongitude(longitude), getRadius(radius), skip, limit)
	if err != nil {
		return nil, err
	}
	return st, nil
}

// UpdateByID update storage
func (s *StationService) UpdateByID(id string, st *models.Station, ownerID string) (*models.Station, error) {
	err := s.storage.Station().UpdateByID(id, st, ownerID)
	if err != nil {
		return nil, err
	}
	st, err = s.storage.Station().FindByID(id)
	if err != nil {
		return nil, err
	}
	return st, nil
}

// DeleteByID delete storage
func (s *StationService) DeleteByID(id string, ownerID string) error {
	return s.storage.Station().DeleteByID(id, ownerID)
}

func (s *StationService) Read(skip int, limit int) ([]models.Station, error) {
	st, err := s.storage.Station().Read(skip, limit)
	if err != nil {
		return nil, err
	}
	return st, nil
}

func getLongitude(lat float64) float64 {
	lat = math.Remainder(lat, 360)
	if lat > 180 {
		lat = -180 + math.Remainder(lat, 180)
	}
	return lat
}

func getLatitude(lng float64) float64 {
	if lng > 90 {
		lng = 90
	}
	if lng < -90 {
		lng = -90
	}
	return lng
}

func getRadius(rad int) float64 {
	if rad > 40074 {
		rad = 40074
	}
	radius := float64(float64(float64(rad)/float64(6378)) * float64(180) / math.Pi)
	return radius
}
