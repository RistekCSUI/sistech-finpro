package thread

import (
	"errors"
	"github.com/RistekCSUI/sistech-finpro/application"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ViewService interface {
		CreateThread(request dto.CreateThreadRequest) (*dto.CreateThreadResponse, error)
		GetAllThreadyByCategory(request dto.GetAllThreadRequest) (*dto.GetAllThreadResponse, error)
		EditThread(request dto.EditThreadRequest) (*dto.EditThreadResponse, error)
		DeleteThread(request dto.DeleteThreadRequest) (*dto.DeleteThreadResponse, error)
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
			ID:        postId.(primitive.ObjectID),
			Content:   request.FirstPost.Content,
			Upvote:    0,
			Downvote:  0,
			Owner:     request.Owner,
			Edited:    false,
			IsStarter: true,
		},
	}

	return response, nil
}

func (v *viewService) GetAllThreadyByCategory(request dto.GetAllThreadRequest) (*dto.GetAllThreadResponse, error) {
	data, category, err := v.application.ThreadService.FindAll(request)
	if err != nil {
		return nil, err
	}

	response := &dto.GetAllThreadResponse{
		Name: category.Name,
		Data: *data,
	}

	return response, nil
}

func (v *viewService) EditThread(request dto.EditThreadRequest) (*dto.EditThreadResponse, error) {
	data, err := v.application.ThreadService.Update(request)
	if err != nil {
		return nil, err
	}

	res := &dto.EditThreadResponse{
		ModifiedCount: data.(int64),
		Name:          request.Name,
	}

	if res.ModifiedCount == 0 {
		return nil, errors.New("failed to edit thread")
	}

	return res, nil
}

func (v *viewService) DeleteThread(request dto.DeleteThreadRequest) (*dto.DeleteThreadResponse, error) {
	data, err := v.application.ThreadService.Delete(request)
	if err != nil {
		return nil, err
	}

	res := &dto.DeleteThreadResponse{
		DeletedCount: data.(int64),
	}

	if res.DeletedCount == 0 {
		return nil, errors.New("failed to delete thread")
	}

	return res, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
