package webtag

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/entity"
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/ghostrepo00/go-dashboard-api/domain/model/exception"
	"gorm.io/gorm"
)

type repository struct {
	dbClient *gorm.DB
}

type WebtagRepository interface {
	GetWebtag() (result []model.Webtag, err error)
	GetWebtagById(id string) (result model.Webtag, err error)
	CreateWebtag(model *model.Webtag) (err error)
	UpdateWebtag(model *model.Webtag) (err error)
	DeleteWebtag(id string) (err error)
}

func toModel(input entity.Webtag) (result model.Webtag) {
	result.Id = input.Id
	result.Url = input.Url
	result.Title = input.Title
	result.Note = input.Note
	for _, tag := range input.Tags {
		result.Tags = append(result.Tags, tag)
	}
	return
}

func toEntity(input *model.Webtag) (result entity.Webtag) {
	result.Id = input.Id
	result.Url = input.Url
	result.Title = input.Title
	result.Note = input.Note
	for _, tag := range input.Tags {
		result.Tags = append(result.Tags, tag)
	}
	return
}

func (repo *repository) GetWebtag() (result []model.Webtag, err error) {
	var dbResult []entity.Webtag
	if err = repo.dbClient.Find(&dbResult).Error; err == nil {
		for _, v := range dbResult {
			result = append(result, toModel(v))
		}
	}
	return
}

func (repo *repository) GetWebtagById(id string) (result model.Webtag, err error) {
	var dbResult *entity.Webtag
	if err = repo.dbClient.First(&dbResult, "id = ?", id).Error; err == nil {
		result = toModel(*dbResult)
	} else {
		return model.Webtag{}, &exception.NotFound{}
	}
	return
}

func (repo *repository) CreateWebtag(model *model.Webtag) (err error) {
	entity := toEntity(model)
	return repo.dbClient.Create(&entity).Error
}

func (repo *repository) UpdateWebtag(model *model.Webtag) (err error) {
	var entity entity.Webtag
	if err = repo.dbClient.First(&entity, "id = ?", model.Id).Error; err == nil {
		entity = toEntity(model)
		return repo.dbClient.Save(entity).Error
	} else {
		return &exception.NotFound{}
	}
}

func (repo *repository) DeleteWebtag(id string) (err error) {
	return repo.dbClient.Delete(&entity.Webtag{Id: id}).Error
}
