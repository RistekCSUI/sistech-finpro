package post

import (
	"github.com/RistekCSUI/sistech-finpro/application"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
)

type (
	ViewService interface {
		GetAllPostByThread(request dto.GetAllPostRequest) (*[]dto.Post, error)
	}
	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) GetAllPostByThread(request dto.GetAllPostRequest) (*[]dto.Post, error) {
	data, err := v.application.PostService.FindAll(request)

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
