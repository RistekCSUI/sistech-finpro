package access

import (
	"context"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Service interface {
		Insert(access dto.CreateAccessDto) (interface{}, error)
		FindByToken(token string) (dto.CreateAccessResponse, error)
	}

	service struct {
		DB *mongo.Collection
	}
)

func (s *service) Insert(access dto.CreateAccessDto) (interface{}, error) {
	acRow := bson.D{{"name", access.Name}}
	res, err := s.DB.InsertOne(context.TODO(), acRow)
	if err != nil {
		return 0, err
	}
	return res.InsertedID, nil
}

func (s *service) FindByToken(token string) (dto.CreateAccessResponse, error) {
	var result = dto.CreateAccessResponse{}

	objectId, err := primitive.ObjectIDFromHex(token)
	if err != nil {
		return result, err
	}

	err = s.DB.FindOne(context.TODO(), bson.D{{"_id", objectId}}).Decode(&result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func NewService(db *mongo.Database) Service {
	return &service{
		DB: db.Collection("access"),
	}
}
