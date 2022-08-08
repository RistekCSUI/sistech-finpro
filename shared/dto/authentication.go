package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	USER  = "user"
	ADMIN = "admin"
)

type RegisterDto struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

type RegisterRequest struct {
	RegisterDto
	Token string
}

type RegisterResponse struct {
	ID       primitive.ObjectID `json:"id"`
	Username string             `json:"username"`
	Role     string             `json:"role"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	Username    string             `bson:"username"`
	AccessToken string             `bson:"accessToken"`
	Password    string             `bson:"password"`
	Role        string             `bson:"role"`
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
	Role     string `json:"role"`
	Token    string `json:"token"`
}
