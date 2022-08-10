package category

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
		Insert(request dto.CreateCategoryRequest) (interface{}, error)
		Update(request dto.EditCategoryRequest) (interface{}, error)
		Delete(request dto.DeleteCategoryRequest) (interface{}, error)
		FindAll(request dto.GetAllCategoryRequest) (*[]dto.Category, error)
	}

	service struct {
		DB *mongo.Collection
	}
)

func (s *service) Insert(request dto.CreateCategoryRequest) (interface{}, error) {
	row := bson.D{
		{"accessToken", request.Token},
		{"name", request.Name},
	}

	exist := s.DB.FindOne(context.TODO(), row)
	if exist.Err() == nil {
		return nil, errors.New("duplicate category name for this access key")
	}

	res, err := s.DB.InsertOne(context.TODO(), row)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

func (s *service) Update(request dto.EditCategoryRequest) (interface{}, error) {
	row := bson.D{
		{"accessToken", request.Token},
		{"name", request.Name},
	}

	exist := s.DB.FindOne(context.TODO(), row)
	if exist.Err() == nil {
		return nil, errors.New("duplicate category name for this access key")
	}

	id, _ := primitive.ObjectIDFromHex(request.ID)
	result, err := s.DB.UpdateOne(
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

func (s *service) Delete(request dto.DeleteCategoryRequest) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(request.ID)
	result, err := s.DB.DeleteOne(context.TODO(), bson.M{"_id": id, "accessToken": request.Token})
	if err != nil {
		return nil, err
	}
	return result.DeletedCount, nil
}

func (s *service) FindAll(request dto.GetAllCategoryRequest) (*[]dto.Category, error) {
	var result []dto.Category

	cur, err := s.DB.Find(context.TODO(), bson.D{{"accessToken", request.Token}})
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem dto.Category
		_ = cur.Decode(&elem)
		result = append(result, elem)
	}

	if err = cur.Err(); err != nil {
		return nil, err
	}

	return &result, nil
}

func NewService(db *mongo.Database) Service {
	return &service{
		DB: db.Collection("category"),
	}
}
