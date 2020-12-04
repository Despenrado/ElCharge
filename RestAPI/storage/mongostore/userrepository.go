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

// UserRepository ...
type UserRepository struct {
	storage *Storage
	col     *mongo.Collection
}

// NewUserRepository constructor
func NewUserRepository(col *mongo.Collection) *UserRepository {
	return &UserRepository{
		col: col,
	}
}

// Create ...
func (r *UserRepository) Create(u *models.User) (string, error) {
	res, err := r.col.InsertOne(context.TODO(), u)
	if err != nil {
		return "", err
	}
	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

// FindByID ...
func (r *UserRepository) FindByID(id string) (*models.User, error) {
	idi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": idi}
	res := r.col.FindOne(context.TODO(), filter)
	u := &models.User{}
	err = res.Decode(u)
	if err != nil {
		return nil, utils.ErrRecordNotFound
	}
	return u, nil
}

// FindByEmail ...
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	filter := bson.M{"email": email}
	res := r.col.FindOne(context.TODO(), filter)
	u := &models.User{}
	err := res.Decode(u)
	if err != nil {
		return nil, utils.ErrRecordNotFound
	}
	return u, nil
}

// UpdateByID ...
func (r *UserRepository) UpdateByID(id string, u *models.User) error {
	idi, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id":             idi,
		"model.update_at": u.UpdateAt,
	}
	update := bson.M{
		"model.update_at": models.GetTimeNow(),
	}
	if u.UserName != "" {
		update["user_name"] = u.UserName
	}
	if u.Email != "" {
		update["email"] = u.Email
	}
	if u.Password != "" {
		update["password"], err = models.EncryptString(u.Password)
		if err != nil {
			return err
		}
	}
	_, err = r.col.UpdateOne(context.TODO(), filter, bson.M{
		"$set": update})
	if err != nil {
		return err
	}
	return nil
}

// DeleteByID ...
func (r *UserRepository) DeleteByID(id string) error {
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

func (r *UserRepository) Read(skip int, limit int) ([]models.User, error) {
	cursor, err := r.col.Aggregate(
		context.TODO(),
		mongo.Pipeline{
			bson.D{{"$skip", skip}},
			bson.D{{"$limit", limit}},
		})
	log.Println("repa")
	if err != nil {
		return nil, err
	}
	users := []models.User{}
	log.Println(users)
	err = cursor.All(context.TODO(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
