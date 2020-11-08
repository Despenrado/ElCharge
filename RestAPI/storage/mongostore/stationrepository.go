package mongostorage

import (
	"context"
	"time"

	"github.com/Despenrado/ElCharge/RestAPI/models"
	"github.com/Despenrado/ElCharge/RestAPI/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// StationRepository struct
type StationRepository struct {
	storage *Storage
	col     *mongo.Collection
}

// NewStationRepository constructor
func NewStationRepository(col *mongo.Collection) *StationRepository {
	return &StationRepository{
		col: col,
	}
}

// Create ...
func (r *StationRepository) Create(s *models.Station) (string, error) {
	res, err := r.col.InsertOne(context.TODO(), s)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

// FindByID find item by id and return it (form map in RAM)
func (r *StationRepository) FindByID(id string) (*models.Station, error) {
	idi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": idi}
	res := r.col.FindOne(context.TODO(), filter)
	s := &models.Station{}
	err = res.Decode(s)
	if err != nil {
		return nil, utils.ErrRecordNotFound
	}
	return s, nil
}

// FindByLocation find item by email and return it (form map in RAM)
func (r *StationRepository) FindByLocation(location string) (*models.Station, error) {
	filter := bson.M{"location": location}
	res := r.col.FindOne(context.TODO(), filter)
	s := &models.Station{}
	err := res.Decode(s)
	if err != nil {
		return nil, utils.ErrRecordNotFound
	}
	return s, nil
}

// UpdateByID update item in db (map in RAM)
func (r *StationRepository) UpdateByID(id string, s *models.Station) error {
	idi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id":             idi,
		"model.update_at": s.UpdateAt,
	}
	update := bson.M{
		"model.update_at": time.Now(),
	}
	if s.Description != "" {
		update["description"] = s.Description
	}
	if s.Raiting != 0 {
		update["raiting"] = s.Raiting
	}
	if s.StationName != "" {
		update["station_name"] = s.StationName
	}
	if s.Location != "" {
		update["location"] = s.Location
	}
	_, err = r.col.UpdateOne(context.TODO(), filter, bson.M{
		"$set": update})
	if err != nil {
		return err
	}
	return nil
}

// DeleteByID delete item by if from db (map in RAM)
func (r *StationRepository) DeleteByID(id string) error {
	idi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": idi}
	res, err := r.col.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return utils.ErrRecordNotFound
	}
	return nil
}

func (r *StationRepository) Read(skip int, limit int) ([]models.Station, error) {
	cursor, err := r.col.Aggregate(
		context.TODO(),
		mongo.Pipeline{
			bson.D{{"$skip", skip}},
			bson.D{{"$limit", limit}},
		})
	if err != nil {
		return nil, err
	}
	stations := []models.Station{}
	err = cursor.All(context.TODO(), &stations)
	if err != nil {
		return nil, err
	}
	return stations, nil
}
