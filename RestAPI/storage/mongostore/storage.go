package mongostorage

import (
	"context"

	"github.com/Despenrado/ElCharge/RestAPI/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Storage ...
type Storage struct {
	userRepository    *UserRepository
	stationRepository *StationRepository
	commentRepository *CommentRepository
}

// NewStorage constructor
func NewStorage(ur *UserRepository, sr *StationRepository, cr *CommentRepository) *Storage {
	s := &Storage{
		userRepository:    ur,
		stationRepository: sr,
		commentRepository: cr,
	}
	s.userRepository.storage = s
	return s
}

// ConfigureRepository configure repository
func ConfigureRepository(db *mongo.Database, colName string) *mongo.Collection {
	col := db.Collection(colName)
	return col
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

// Station ...
func (s *Storage) Station() storage.StationRepository {
	if s.stationRepository != nil {
		return s.stationRepository
	}
	s.stationRepository = &StationRepository{
		storage: s,
	}
	return s.stationRepository
}

// Comment ...
func (s *Storage) Comment() storage.CommentRepository {
	if s.commentRepository != nil {
		return s.commentRepository
	}
	s.commentRepository = &CommentRepository{
		storage: s,
	}
	return s.commentRepository
}
