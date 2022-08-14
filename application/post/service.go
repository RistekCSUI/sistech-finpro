package post

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
		FindAll(request dto.GetAllPostRequest) (*[]dto.Post, error)
		Insert(request dto.CreatePostRequest) (interface{}, error)
		Vote(request dto.CreateVoteRequest) (interface{}, *dto.Post, error)
		Update(request dto.EditPostRequest) (interface{}, error)
		Delete(request dto.DeletePostRequest) (interface{}, error)
	}
	service struct {
		DB     *mongo.Collection
		Thread *mongo.Collection
		Post   *mongo.Collection
	}
)

func (s *service) Insert(request dto.CreatePostRequest) (interface{}, error) {
	// Check if thread exist
	id, _ := primitive.ObjectIDFromHex(request.ThreadID)
	thread := bson.D{
		{"_id", id},
		{"accessToken", request.Token},
	}
	exist := s.Thread.FindOne(context.TODO(), thread)
	if exist.Err() != nil {
		return nil, errors.New("no thread for given id")
	}

	post := bson.D{
		{"accessToken", request.Token},
		{"content", request.Content},
		{"threadId", request.ThreadID},
		{"replyId", request.ReplyID},
		{"upvote", 0},
		{"downvote", 0},
		{"edited", false},
		{"owner", request.Owner},
		{"isStarter", false},
	}

	res, err := s.DB.InsertOne(context.TODO(), post)
	if err != nil {
		return nil, err
	}

	return res.InsertedID, nil
}

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

func (s *service) Vote(request dto.CreateVoteRequest) (interface{}, *dto.Post, error) {
	id, _ := primitive.ObjectIDFromHex(request.PostID)
	row := bson.D{
		{"_id", id},
		{"accessToken", request.Token},
	}

	var post dto.Post
	err := s.DB.FindOne(context.TODO(), row).Decode(&post)
	if err != nil {
		return nil, nil, errors.New("no post for given id")
	}

	if post.Owner == request.RequesterID {
		return nil, nil, errors.New("cant vote your own post")
	}

	newVote := bson.D{
		{dto.UPVOTE, post.Upvote + 1},
	}
	post.Upvote += 1

	if request.VoteType == dto.DOWNVOTE {
		newVote = bson.D{{dto.DOWNVOTE, post.Downvote + 1}}
		post.Upvote -= 1
		post.Downvote += 1
	}

	result, err := s.DB.UpdateOne(
		context.TODO(),
		bson.M{"_id": id, "accessToken": request.Token},
		bson.D{
			{"$set", newVote},
		},
	)

	if err != nil {
		return nil, nil, err
	}

	return result.ModifiedCount, &post, nil
}

func (s *service) Delete(request dto.DeletePostRequest) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(request.PostID)
	result, err := s.DB.DeleteOne(context.TODO(), bson.M{
		"_id":         id,
		"accessToken": request.Token,
	})
	if err != nil {
		return nil, err
	}
	return result.DeletedCount, nil
}

func (s *service) Update(request dto.EditPostRequest) (interface{}, error) {
	id, _ := primitive.ObjectIDFromHex(request.PostID)
	result, err := s.DB.UpdateOne(
		context.TODO(),
		bson.M{"_id": id, "accessToken": request.Token, "owner": request.OwnerID},
		bson.D{
			{"$set", bson.D{
				{"content", request.Content},
				{"edited", true},
			}},
		},
	)

	if err != nil {
		return nil, err
	}

	return result.ModifiedCount, nil
}

func NewService(db *mongo.Database) Service {
	return &service{
		DB:     db.Collection("post"),
		Thread: db.Collection("thread"),
		Post:   db.Collection("post"),
	}
}
