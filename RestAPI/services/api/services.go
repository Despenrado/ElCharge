package api

import "github.com/Despenrado/ElCharge/RestAPI/models"

type UserService interface {
	CreateUser(*models.User) (*models.User, error)
	Login(*models.User) (*models.User, error)
	FindByID(string) (*models.User, error)
	UpdateByID(string, *models.User) (*models.User, error)
	DeleteByID(id string) error
	Read(int, int) ([]models.User, error)
}

type StationService interface {
	CreateStation(*models.Station) (*models.Station, error)
	FindByID(string) (*models.Station, error)
	FindByLocation(float64, float64) (*models.Station, error)
	FindByDescription(string) ([]models.Station, error)
	FindByName(string) ([]models.Station, error)
	FindInRadius(float64, float64, int, int, int) ([]models.Station, error)
	UpdateByID(string, *models.Station, string) (*models.Station, error)
	DeleteByID(string, string) error
	Read(int, int) ([]models.Station, error)
}

type CommentService interface {
	CreateComment(string, *models.Comment) (*models.Comment, error)
	FindByID(string, string) (*models.Comment, error)
	FindByUserName(string, string) (*models.Comment, error)
	UpdateByID(string, string, *models.Comment) (*models.Comment, error)
	DeleteByID(string, string) error
	Read(string, int, int) ([]models.Comment, error)
}
