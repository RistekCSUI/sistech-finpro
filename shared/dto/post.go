package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type PostInThread struct {
	Content string `json:"content" validate:"required"`
}

type Post struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Content   string             `json:"content" bson:"content"`
	Downvote  int64              `json:"downvote" bson:"downvote"`
	Upvote    int64              `json:"upvote" bson:"upvote"`
	ReplyID   string             `json:"replyId" bson:"replyId"`
	Owner     string             `json:"owner" bson:"owner"`
	IsStarter bool               `json:"isStarter" bson:"isStarter"`
	Edited    bool               `json:"edited" bson:"edited"`
}

type CreatePostResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Content   string             `json:"content"`
	Upvote    int                `json:"upvote"`
	Downvote  int                `json:"downvote"`
	Owner     string             `json:"owner"`
	Edited    bool               `json:"edited"`
	IsStarter bool               `json:"isStarter"`
}

type GetAllPostRequest struct {
	ThreadID string
	Token    string
}
