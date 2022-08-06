package access

import (
	"github.com/RistekCSUI/sistech-finpro/application"
	"github.com/RistekCSUI/sistech-finpro/shared"
	"github.com/RistekCSUI/sistech-finpro/shared/dto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	ViewService interface {
		Register(body dto.CreateAccessDto) (dto.CreateAccessResponse, error)
	}

	viewService struct {
		application application.Holder
		shared      shared.Holder
	}
)

func (v *viewService) Register(body dto.CreateAccessDto) (dto.CreateAccessResponse, error) {
	var (
		response dto.CreateAccessResponse
	)

	v.shared.Logger.Info("insert access into db for name: ", body.Name)
	id, err := v.application.AccessService.Insert(body)
	if err != nil {
		v.shared.Logger.Errorf("error inserting access to db for name: %s , with error: %s", body.Name, err.Error())
		return response, err
	}

	response = dto.CreateAccessResponse{
		ID:   id.(primitive.ObjectID),
		Name: body.Name,
	}

	return response, nil
}

func NewViewService(application application.Holder, shared shared.Holder) ViewService {
	return &viewService{
		application: application,
		shared:      shared,
	}
}
