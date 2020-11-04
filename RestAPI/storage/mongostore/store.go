package mongostorage

import (
	"context"

	"github.com/Despenrado/ElCharge/RestAPI/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Storage ...
type Storage struct {
	userRepository *UserRepository
}

// NewStorage constructor
func NewStorage(ur *UserRepository) *Storage {
	s := &Storage{
		userRepository: ur,
	}
	s.userRepository.storage = s
	return s
}

// ConnectToDB ...
func ConnectToDB(uri string, dbName string) (*mongo.Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(dbName)
	return db, nil
}

// User ...
func (s *Storage) User() storage.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}
	s.userRepository = &UserRepository{
		storage: s,
	}
	return s.userRepository
}
