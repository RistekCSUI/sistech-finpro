package post

import (
	"context"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Service interface {
		FindAll(request dto.GetAllPostRequest) (*[]dto.Post, error)
	}
	service struct {
		DB *mongo.Collection
	}
)

func (s *service) FindAll(request dto.GetAllPostRequest) (*[]dto.Post, error) {
	var result []dto.Post

	filter := bson.D{
		{"accessToken", request.Token},
		{"threadId", request.ThreadID},
	}

	cur, err := s.DB.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	for cur.Next(context.TODO()) {
		var elem dto.Post
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
		DB: db.Collection("post"),
	}
}
