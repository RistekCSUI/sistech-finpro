package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Content string `json:"content" validate:"required"`
}

type CreatePostResponse struct {
	ID      primitive.ObjectID `json:"id"`
	Content string             `json:"content"`
}
