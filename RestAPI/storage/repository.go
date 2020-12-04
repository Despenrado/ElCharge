package storage

import "github.com/Despenrado/ElCharge/RestAPI/models"

// UserRepository ...
type UserRepository interface {
	Create(*models.User) (string, error)
	FindByID(string) (*models.User, error)
	FindByEmail(string) (*models.User, error)
	UpdateByID(string, *models.User) error
	DeleteByID(string) error
	Read(int, int) ([]models.User, error)
}

// StationRepository ...
type StationRepository interface {
	Create(*models.Station) (string, error)
	FindByID(string) (*models.Station, error)
	FindByLocation(float64, float64) (*models.Station, error)
	FindByName(string) ([]models.Station, error)
	FindByDescription(string) ([]models.Station, error)
	FindInRadius(float64, float64, float64, int, int) ([]models.Station, error)
	UpdateByID(string, *models.Station, string) error
	DeleteByID(string, string) error
	Read(int, int) ([]models.Station, error)
	UpdateRaitingByID(string) error
}

// CommentRepository ...
type CommentRepository interface {
	Create(string, *models.Comment) (string, error)
	FindByID(string, string) (*models.Comment, error)
	FindByUserName(string, string) (*models.Comment, error)
	UpdateByID(string, string, *models.Comment) error
	DeleteByID(string, string) error
}
