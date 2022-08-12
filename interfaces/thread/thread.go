package thread

import (
	"github.com/RistekCSUI/sistech-finpro/application"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ViewService interface {
		CreateThread(request dto.CreateThreadRequest) (*dto.CreateThreadResponse, error)
		GetAllThreadyByCategory(request dto.GetAllThreadRequest) (*[]dto.Thread, error)
	}
	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) CreateThread(request dto.CreateThreadRequest) (*dto.CreateThreadResponse, error) {
	threadId, postId, err := v.application.ThreadService.Insert(request)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateThreadResponse{
		ID:   threadId.(primitive.ObjectID),
		Name: request.Name,
		FirstPost: dto.CreatePostResponse{
			ID:      postId.(primitive.ObjectID),
			Content: request.FirstPost.Content,
		},
	}

	return response, nil
}

func (v *viewService) GetAllThreadyByCategory(request dto.GetAllThreadRequest) (*[]dto.Thread, error) {
	data, err := v.application.ThreadService.FindAll(request)
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
