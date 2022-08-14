package thread

import (
	"context"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Service interface {
		Insert(request dto.CreateThreadRequest) (interface{}, interface{}, error)
		FindAll(request dto.GetAllThreadRequest) (*[]dto.Thread, *dto.Category, error)
		Update(request dto.EditThreadRequest) (interface{}, error)
		Delete(request dto.DeleteThreadRequest) (interface{}, error)
	}

	service struct {
		Post     *mongo.Collection
		Thread   *mongo.Collection
		Category *mongo.Collection
	}
)

func (s *service) Insert(request dto.CreateThreadRequest) (interface{}, interface{}, error) {
	thread := bson.D{
		{"accessToken", request.Token},
		{"name", request.Name},
		{"categoryId", request.CategoryID},
	}

	// Check if there is existed a category for given id and accessToken
	id, _ := primitive.ObjectIDFromHex(request.CategoryID)
	category := bson.D{
		{"_id", id},
		{"accessToken", request.Token},
	}

	exist := s.Category.FindOne(context.TODO(), category)
	if exist.Err() != nil {
		return nil, nil, errors.New("no category for given id")
	}

	// Check for duplicate thread name for the same category
	threadExist := s.Thread.FindOne(context.TODO(), thread)
	if threadExist.Err() == nil {
		return nil, nil, errors.New("already exist a thread with given name")
	}

	res, err := s.Thread.InsertOne(context.TODO(), thread)
	if err != nil {
		return nil, nil, err
	}

	post := bson.D{
		{"accessToken", request.Token},
		{"content", request.FirstPost.Content},
		{"threadId", res.InsertedID.(primitive.ObjectID).Hex()},
		{"replyId", ""},
		{"upvote", 0},
		{"downvote", 0},
		{"edited", false},
		{"owner", request.Owner},
		{"isStarter", true},
	}

	resPost, err := s.Post.InsertOne(context.TODO(), post)
	if err != nil {
		return nil, nil, err
	}

	return res.InsertedID, resPost.InsertedID, nil
}

func (s *service) FindAll(request dto.GetAllThreadRequest) (*[]dto.Thread, *dto.Category, error) {
	var result []dto.Thread
	var category dto.Category

	id, _ := primitive.ObjectIDFromHex(request.CategoryID)
	err := s.Category.FindOne(context.TODO(), bson.D{
		{"_id", id},
		{"accessToken", request.Token},
	}).Decode(&category)
	if err != nil {
		return nil, nil, errors.New("no category found for given id")
	}

	filter := bson.D{
		{"accessToken", request.Token},
		{"categoryId", request.CategoryID},
	}

	cur, err := s.Thread.Find(context.TODO(), filter)
	if err != nil {
		return nil, nil, err
	}

	for cur.Next(context.TODO()) {
		var elem dto.Thread
		_ = cur.Decode(&elem)
		result = append(result, elem)
	}

	if err = cur.Err(); err != nil {
		return nil, nil, err
	}

	return &result, &category, nil
}

func (s *service) Update(request dto.EditThreadRequest) (interface{}, error) {
	row := bson.D{
		{"accessToken", request.Token},
		{"name", request.Name},
	}

	var exist dto.Thread
	err := s.Thread.FindOne(context.TODO(), row).Decode(&exist)
	if err == nil && exist.Name != request.Name {
		return nil, errors.New("duplicate thread name for this access key")
	}

	id, _ := primitive.ObjectIDFromHex(request.ThreadID)
	result, err := s.Thread.UpdateOne(
		context.TODO(),
		bson.M{"_id": id, "accessToken": request.Token},
		bson.D{
			{"$set", bson.D{{"name", request.Name}}},
		},
	)

	if err != nil {
		return nil, err
	}

	return result.ModifiedCount, nil
}

func (s *service) Delete(request dto.DeleteThreadRequest) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(request.ID)
	result, err := s.Thread.DeleteOne(context.TODO(), bson.M{"_id": id, "accessToken": request.Token})
	if err != nil {
		return nil, err
	}
	_, _ = s.Post.DeleteMany(context.TODO(), bson.D{{"threadId", request.ID}})
	return result.DeletedCount, nil
}

func NewService(db *mongo.Database) Service {
	return &service{
		Post:     db.Collection("post"),
		Thread:   db.Collection("thread"),
		Category: db.Collection("category"),
	}
}
