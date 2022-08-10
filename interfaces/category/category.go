package category

import (
	"github.com/RistekCSUI/sistech-finpro/application"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ViewService interface {
		CreateCategory(request dto.CreateCategoryRequest) (*dto.Category, error)
		EditCategory(request dto.EditCategoryRequest) (*dto.EditCategoryResponse, error)
	}

	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateCategory(request dto.CreateCategoryRequest) (*dto.Category, error) {
	data, err := v.application.CategoryService.Insert(request)
	if err != nil {
		return nil, err
	}

	res := &dto.Category{
		ID:   data.(primitive.ObjectID),
		Name: request.Name,
	}
	return res, nil
}

func (v *viewService) EditCategory(request dto.EditCategoryRequest) (*dto.EditCategoryResponse, error) {
	data, err := v.application.CategoryService.Update(request)
	if err != nil {
		return nil, err
	}

	res := &dto.EditCategoryResponse{
		ModifiedCount: data.(int64),
		Name:          request.Name,
	}

	return res, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
