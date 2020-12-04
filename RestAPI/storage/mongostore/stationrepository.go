package mongostorage

import (
	"context"
	"log"

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
func (r *StationRepository) FindByLocation(latitude float64, longitude float64) (*models.Station, error) {
	filter := bson.M{"latitude": latitude, "longitude": longitude}
	res := r.col.FindOne(context.TODO(), filter)
	s := &models.Station{}
	err := res.Decode(s)
	if err != nil {
		return nil, utils.ErrRecordNotFound
	}
	return s, nil
}

func (r *StationRepository) FindByName(name string) ([]models.Station, error) {
	filter := bson.D{{"station_name", primitive.Regex{Pattern: name, Options: "i"}}}
	cursor, err := r.col.Aggregate(
		context.TODO(),
		mongo.Pipeline{
			bson.D{{"$match",
				filter}},
		})
	if err != nil {
		return nil, err
	}
	s := []models.Station{}
	err = cursor.All(context.TODO(), &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *StationRepository) FindByDescription(text string) ([]models.Station, error) {
	filter := bson.D{{"description", primitive.Regex{Pattern: text, Options: "i"}}}
	log.Println(filter)
	cursor, err := r.col.Aggregate(
		context.TODO(),
		mongo.Pipeline{
			bson.D{{"$match",
				filter}},
		})
	if err != nil {
		return nil, err
	}
	s := []models.Station{}
	err = cursor.All(context.TODO(), &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *StationRepository) FindInRadius(latitude float64, longitude float64, radius float64, skip int, limit int) ([]models.Station, error) {
	matchInRadius := bson.D{{"$match", bson.M{
		"$and": []interface{}{
			bson.M{"latitude": bson.M{
				"$gte": latitude - radius}},
			bson.M{"latitude": bson.M{
				"$lte": latitude + radius}},
			bson.M{"longitude": bson.M{
				"$gte": longitude - radius}},
			bson.M{"longitude": bson.M{
				"$lte": longitude + radius}},
		}},
	}}
	cursor, err := r.col.Aggregate(
		context.TODO(),
		mongo.Pipeline{
			matchInRadius,
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

// UpdateByID update item in db (map in RAM)
func (r *StationRepository) UpdateByID(id string, s *models.Station, ownerID string) error {
	idi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id":             idi,
		"model.update_at": s.UpdateAt,
		"owner_id":        ownerID,
	}
	update := bson.M{
		"model.update_at": models.GetTimeNow(),
	}
	if s.Description != "" {
		update["description"] = s.Description
	}
	if s.Rating != 0 {
		update["raiting"] = s.Rating
	}
	if s.StationName != "" {
		update["station_name"] = s.StationName
	}
	_, err = r.col.UpdateOne(context.TODO(), filter, bson.M{
		"$set": update})
	if err != nil {
		return err
	}
	return nil
}

// DeleteByID delete item by if from db (map in RAM)
func (r *StationRepository) DeleteByID(id string, ownerID string) error {
	idi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": idi, "owner_id": ownerID}
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

func (r *StationRepository) UpdateRaitingByID(id string) error {
	idi, err := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": idi}
	avg := bson.D{{
		"$set", bson.M{
			"rating": bson.M{
				"$divide": []interface{}{
					bson.M{
						"$reduce": bson.M{
							"input":        "$comments",
							"initialValue": 0,
							"in":           bson.M{"$add": []interface{}{"$$value", "$$this.rating"}},
						},
					},
					bson.M{
						"$cond": []interface{}{
							bson.M{"$ne": []interface{}{bson.M{"$size": "$comments"}, 0}},
							bson.M{"$size": "$comments"},
							1,
						},
					},
				},
			},
		},
	}}
	_, err = r.col.UpdateOne(
		context.TODO(),
		filter,
		mongo.Pipeline{avg})
	return err
}
