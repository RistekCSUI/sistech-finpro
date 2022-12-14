package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	UPVOTE   = "upvote"
	DOWNVOTE = "downvote"
)

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

type CreatePostDto struct {
	ThreadID string `json:"threadId" validate:"required"`
	Content  string `json:"content" validate:"required"`
	ReplyID  string `json:"replyId"`
}

type CreatePostRequest struct {
	Token    string
	Content  string
	Owner    string
	ThreadID string
	ReplyID  string
}

type CreatePostResponse struct {
	ID        primitive.ObjectID `json:"id"`
	Content   string             `json:"content"`
	Upvote    int                `json:"upvote"`
	Downvote  int                `json:"downvote"`
	Owner     string             `json:"owner"`
	Edited    bool               `json:"edited"`
	IsStarter bool               `json:"isStarter"`
	ReplyID   ReplyPost          `json:"reply"`
}

type ReplyPost struct {
	ID      primitive.ObjectID `json:"id"`
	Content string             `json:"content"`
}

type GetAllPostRequest struct {
	ThreadID string
	Token    string
}

type GetAllPostResponse struct {
	Name string `json:"name"`
	Data []Post `json:"data"`
}

type VoteDto struct {
	VoteType string `json:"voteType"`
	PostID   string `json:"postId"`
}

type CreateVoteRequest struct {
	Token       string
	PostID      string
	VoteType    string
	RequesterID string
}

type CreateVoteResponse struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Upvote        int64 `json:"upvote"`
	Downvote      int64 `json:"downvote"`
}

type EditPostDto struct {
	Content string `json:"content"`
}

type EditPostRequest struct {
	Token   string
	PostID  string
	Content string
	OwnerID string
}

type EditPostResponse struct {
	Content       string `json:"content"`
	ModifiedCount int64  `json:"modifiedCount"`
}

type DeletePostRequest struct {
	Token       string
	PostID      string
	RequesterID string
}

type DeletePostResponse struct {
	DeleteCount int64 `json:"deleteCount"`
}
