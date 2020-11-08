package v1

import (
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
	_, err := s.storage.Station().FindByLocation(st.Location)
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

func (s *StationService) FindByLocation(location string) (*models.Station, error) {
	st, err := s.storage.Station().FindByLocation(location)
	if err != nil {
		return nil, err
	}
	return st, nil
}

// UpdateByID update storage
func (s *StationService) UpdateByID(id string, st *models.Station) (*models.Station, error) {
	err := s.storage.Station().UpdateByID(id, st)
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
func (s *StationService) DeleteByID(id string) error {
	return s.storage.Station().DeleteByID(id)
}

func (s *StationService) Read(skip int, limit int) ([]models.Station, error) {
	st, err := s.storage.Station().Read(skip, limit)
	if err != nil {
		return nil, err
	}
	return st, nil
}
