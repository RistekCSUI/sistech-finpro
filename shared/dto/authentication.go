package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type RegisterDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	RegisterDto
	Token string
}

type RegisterResponse struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username"`
	AccessToken string             `bson:"accessToken"`
	Password    string             `bson:"password"`
}

type LoginDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginRequest struct {
	LoginDto
	AccessToken string
}

type LoginResponse struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}
