package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateThreadDto struct {
	CategoryID string `json:"categoryId" validate:"required"`
	Name       string `json:"name" validate:"required"`
	FirstPost  Post   `json:"firstPost" validate:"required"`
}

type CreateThreadRequest struct {
	CategoryID string
	Name       string
	FirstPost  Post
	Token      string
}

type CreateThreadResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Name      string             `json:"name"`
	FirstPost CreatePostResponse `json:"firstPost"`
}
