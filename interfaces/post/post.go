package post

import (
	"github.com/RistekCSUI/sistech-finpro/application"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ViewService interface {
		GetAllPostByThread(request dto.GetAllPostRequest) (*[]dto.Post, error)
		CreatePost(request dto.CreatePostRequest) (*dto.CreatePostResponse, error)
		VotePost(request dto.CreateVoteRequest) (*dto.CreateVoteResponse, error)
		EditPost(request dto.EditPostRequest) (*dto.EditPostResponse, error)
		DeletePost(request dto.DeletePostRequest) (*dto.DeletePostResponse, error)
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

func (v *viewService) CreatePost(request dto.CreatePostRequest) (*dto.CreatePostResponse, error) {
	id, err := v.application.PostService.Insert(request)
	if err != nil {
		return nil, err
	}

	response := dto.CreatePostResponse{
		ID:        id.(primitive.ObjectID),
		Content:   request.Content,
		Upvote:    0,
		Downvote:  0,
		Owner:     request.Owner,
		Edited:    false,
		IsStarter: false,
		ReplyID:   request.ReplyID,
	}

	return &response, nil
}

func (v *viewService) VotePost(request dto.CreateVoteRequest) (*dto.CreateVoteResponse, error) {
	count, post, err := v.application.PostService.Vote(request)
	if err != nil {
		return nil, err
	}

	response := &dto.CreateVoteResponse{
		Upvote:        post.Upvote,
		Downvote:      post.Downvote,
		ModifiedCount: count.(int64),
	}

	return response, nil
}

func (v *viewService) EditPost(request dto.EditPostRequest) (*dto.EditPostResponse, error) {
	count, err := v.application.PostService.Update(request)
	if err != nil {
		return nil, err
	}

	response := &dto.EditPostResponse{
		ModifiedCount: count.(int64),
		Content:       request.Content,
	}

	if response.ModifiedCount == 0 {
		return nil, errors.New("failed to edit post")
	}

	return response, nil
}

func (v *viewService) DeletePost(request dto.DeletePostRequest) (*dto.DeletePostResponse, error) {
	count, err := v.application.PostService.Delete(request)
	if err != nil {
		return nil, err
	}
	response := &dto.DeletePostResponse{
		DeleteCount: count.(int64),
	}
	if response.DeleteCount == 0 {
		return nil, errors.New("failed to delete post")
	}
	return response, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
