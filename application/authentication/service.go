package authentication

import (
	"context"
	"errors"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Service interface {
		Insert(request dto.RegisterRequest) (interface{}, error)
		FindUser(token string, username string) (dto.User, error)
		FindUserByID(id string) (dto.User, error)
	}
	service struct {
		DB *mongo.Collection
	}
)

func (s *service) Insert(request dto.RegisterRequest) (interface{}, error) {
	row := bson.D{
		{"username", request.Username},
		{"accessToken", request.Token},
		{"password", request.Password},
		{"role", request.Role},
	}
	exist := s.DB.FindOne(
		context.TODO(),
		bson.D{
			{"username", request.Username},
			{"accessToken", request.Token},
		})

	if exist.Err() == nil {
		return nil, errors.New("duplicate username for this access key")
	}

	res, err := s.DB.InsertOne(context.TODO(), row)
	if err != nil {
		return 0, err
	}
	return res.InsertedID, nil
}

func (s *service) FindUser(token string, username string) (dto.User, error) {
	var result = dto.User{}

	err := s.DB.FindOne(
		context.TODO(),
		bson.D{
			{"username", username},
			{"accessToken", token},
		}).Decode(&result)

	if err != nil {
		return result, err
	}

	return result, nil
}

func (s *service) FindUserByID(id string) (dto.User, error) {
	var result = dto.User{}

	objectId, err := primitive.ObjectIDFromHex(id)
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
		DB: db.Collection("user"),
	}
}
