package webtag

import (
	model "github.com/ghostrepo00/go-dashboard-api/domain/model"
)

type service struct {
	repository WebtagRepository
}

type WebtagService interface {
	GetWebtag() ([]model.Webtag, error)
	GetWebtagById(id string) (model.Webtag, error)
	CreateWebtag(model *model.Webtag) error
	UpdateWebtag(model *model.Webtag) error
	DeleteWebtag(id string) error
}

func (svc *service) GetWebtag() ([]model.Webtag, error) {
	return svc.repository.GetWebtag()
}

func (svc *service) GetWebtagById(id string) (model.Webtag, error) {
	return svc.repository.GetWebtagById(id)
}

func (svc *service) CreateWebtag(model *model.Webtag) error {
	return svc.repository.CreateWebtag(model)
}

func (svc *service) UpdateWebtag(model *model.Webtag) error {
	return svc.repository.UpdateWebtag(model)
}

func (svc *service) DeleteWebtag(id string) error {
	return svc.repository.DeleteWebtag(id)
}
