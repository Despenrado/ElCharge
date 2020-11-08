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

// CommentRepository struct
type CommentRepository struct {
	storage           *Storage
	stationRepository *StationRepository
}

// NewCommentRepository constructor
func NewCommentRepository(sr *StationRepository) *CommentRepository {
	return &CommentRepository{
		stationRepository: sr,
	}
}

// Create create and save item in database
func (r *CommentRepository) Create(sid string, c *models.Comment) (string, error) {
	sidi, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		return "", err
	}
	filter := bson.M{
		"_id": sidi,
	}
	update1 := bson.M{
		"comments": bson.M{
			"$each":     bson.A{c},
			"$position": 0,
		}}
	update2 := bson.M{"model.update_at": time.Now()}
	_, err = r.stationRepository.col.UpdateOne(
		context.TODO(),
		filter,
		bson.M{
			"$push": update1,
			"$set":  update2,
		})
	if err != nil {
		return "", err
	}

	return c.ID, nil
}

// FindByID find item by id and return it
func (r *CommentRepository) FindByID(sid string, id string) (*models.Comment, error) {
	sido, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		return nil, err
	}
	matchSID := bson.D{{
		"$match", bson.M{
			"_id": sido,
		}}}
	unwindComments := bson.D{{"$unwind", "$comments"}}
	raplaceRootComm := bson.D{{
		"$replaceRoot", bson.M{
			"newRoot": "$comments",
		}}}
	machID := bson.D{{
		"$match", bson.M{
			"_id": id,
		}}}
	cursor, err := r.stationRepository.col.Aggregate(
		context.TODO(),
		mongo.Pipeline{
			matchSID,
			unwindComments,
			raplaceRootComm,
			machID,
		})
	if err != nil {
		return nil, err
	}
	comms := []models.Comment{}
	err = cursor.All(context.TODO(), &comms)
	if err != nil {
		return nil, err
	}
	if len(comms) <= 0 {
		return nil, utils.ErrRecordNotFound
	}
	return &comms[0], nil
}

// FindByUserName find item by user name and return it
func (r *CommentRepository) FindByUserName(sid string, uname string) (*models.Comment, error) {
	sido, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		return nil, err
	}
	matchSID := bson.D{{
		"$match", bson.M{
			"_id": sido,
		}}}
	unwindComments := bson.D{{"$unwind", "$comments"}}
	raplaceRootComm := bson.D{{
		"$replaceRoot", bson.M{
			"newRoot": "$comments",
		}}}
	machUName := bson.D{{
		"$match", bson.M{
			"user_name": uname,
		}}}
	cursor, err := r.stationRepository.col.Aggregate(
		context.TODO(),
		mongo.Pipeline{
			matchSID,
			unwindComments,
			raplaceRootComm,
			machUName,
		})
	if err != nil {
		return nil, err
	}
	comms := []models.Comment{}
	err = cursor.All(context.TODO(), &comms)
	if err != nil {
		return nil, err
	}
	if len(comms) <= 0 {
		return nil, utils.ErrRecordNotFound
	}
	return &comms[0], nil
}

// UpdateByID update item in db
func (r *CommentRepository) UpdateByID(sid string, id string, c *models.Comment) error {
	sido, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		return err
	}
	filter := bson.M{
		"_id": sido,
		"comments": bson.M{
			"$elemMatch": bson.M{
				"_id":             id,
				"model.update_at": c.UpdateAt,
			}}}
	update := bson.M{"comments.$.update_at": time.Now()}
	if c.Text != "" {
		update["comments.$.text"] = c.Text
	}
	if c.Raiting != 0 {
		update["comments.$.raiting"] = c.Raiting
	}
	_, err = r.stationRepository.col.UpdateOne(
		context.TODO(),
		filter, bson.M{
			"$set": update,
		})
	if err != nil {
		return err
	}
	return nil
}

// DeleteByID delete item by if from db
func (r *CommentRepository) DeleteByID(sid string, id string) error {
	sido, err := primitive.ObjectIDFromHex(sid)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": sido}
	update := bson.M{
		"comments": bson.M{
			"_id": id,
		}}
	_, err = r.stationRepository.col.UpdateOne(
		context.TODO(),
		filter, bson.M{
			"$pull": update,
		})
	if err != nil {
		return err
	}
	return nil
}
