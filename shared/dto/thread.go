package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Thread struct {
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

type CreateThreadDto struct {
	CategoryID string       `json:"categoryId" validate:"required"`
	Name       string       `json:"name" validate:"required"`
	FirstPost  PostInThread `json:"firstPost" validate:"required"`
}

type CreateThreadRequest struct {
	CategoryID string
	Name       string
	FirstPost  PostInThread
	Token      string
	Owner      string
}

type CreateThreadResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	FirstPost CreatePostResponse `json:"firstPost"`
}

type GetAllThreadRequest struct {
	CategoryID string
	Token      string
}

type EditThreadDto struct {
	Name string `json:"name"`
}

type EditThreadRequest struct {
	EditThreadDto
	ThreadID string
	Token    string
}

type EditThreadResponse struct {
	ModifiedCount int64  `json:"modifiedCount"`
	Name          string `json:"name"`
}

type DeleteThreadRequest struct {
	ID    string
	Token string
}

type DeleteThreadResponse struct {
	DeletedCount int64 `json:"deletedCount"`
}
