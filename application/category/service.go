package category

import (
	"context"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Service interface {
		Insert(request dto.CreateCategoryRequest) (interface{}, error)
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

func NewService(db *mongo.Database) Service {
	return &service{
		DB: db.Collection("category"),
	}
}
