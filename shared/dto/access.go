package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type CreateAccessDto struct {
	Name string `json:"name" validate:"required"`
}

type CreateAccessResponse struct {
	ID   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
}
