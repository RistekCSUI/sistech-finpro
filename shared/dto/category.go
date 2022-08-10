package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type Category struct {
	ID   primitive.ObjectID `json:"id" bson:_id"`
	Name string             `json:"name" bson:"name"`
}

type GetAllCategoryResponse struct {
	Categories []Category `json:"categories"`
}

type CreateCategoryDto struct {
	Name string `json:"name" validate:"required"`
}

type CreateCategoryRequest struct {
	CreateCategoryDto
	Token string
}

type EditCategoryDto struct {
	Name string `json:"name" validate:"required"`
}

type EditCategoryRequest struct {
	EditCategoryDto
	ID    string
	Token string
}

type EditCategoryResponse struct {
	ModifiedCount int64  `json:"modifiedCount"`
	Name          string `json:"name"`
}
