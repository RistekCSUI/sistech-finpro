package category

import (
	"github.com/RistekCSUI/sistech-finpro/application"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ViewService interface {
		CreateCategory(request dto.CreateCategoryRequest) (*dto.Category, error)
		EditCategory(request dto.EditCategoryRequest) (*dto.EditCategoryResponse, error)
		DeleteCategory(request dto.DeleteCategoryRequest) (*dto.DeleteCategoryResponse, error)
		GetAllCategory(request dto.GetAllCategoryRequest) (*[]dto.Category, error)
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

	if res.ModifiedCount == 0 {
		return nil, errors.New("failed to update category")
	}

	return res, nil
}

func (v *viewService) DeleteCategory(request dto.DeleteCategoryRequest) (*dto.DeleteCategoryResponse, error) {
	data, err := v.application.CategoryService.Delete(request)
	if err != nil {
		return nil, err
	}

	res := &dto.DeleteCategoryResponse{
		DeletedCount: data.(int64),
	}

	if res.DeletedCount == 0 {
		return nil, errors.New("failed to delete category")
	}

	return res, nil
}

func (v *viewService) GetAllCategory(request dto.GetAllCategoryRequest) (*[]dto.Category, error) {
	data, err := v.application.CategoryService.FindAll(request)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
