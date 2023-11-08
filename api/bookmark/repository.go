package bookmark

import (
	b64 "encoding/base64"

	"github.com/ghostrepo00/go-dashboard-api/domain/entity"
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/ghostrepo00/go-dashboard-api/domain/model/exception"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	dbClient *gorm.DB
}

type BookmarkRepository interface {
	GetBookmarks() (result []model.Bookmark, err error)
	CreateBookmark(model *model.Bookmark) (err error)
	UpdateBookmark(model *model.Bookmark) (err error)
	DeleteBookmark(id uuid.UUID) (err error)
}

func toEntity(model *model.Bookmark) (result entity.Bookmark, err error) {
	if model.Id != nil {
		result.Id = model.Id
	}
	result.Name = model.Name
	result.URL = model.URL
	result.Icon, err = b64.StdEncoding.DecodeString(model.Icon)
	return
}

func toModel(entity entity.Bookmark) (result model.Bookmark, err error) {
	if entity.Id != nil {
		result.Id = entity.Id
	}
	result.Name = entity.Name
	result.URL = entity.URL
	result.Icon = b64.StdEncoding.EncodeToString(entity.Icon)
	return
}

func (repo *repository) GetBookmarks() (result []model.Bookmark, err error) {
	var dbResult []entity.Bookmark
	if err = repo.dbClient.Order("created_at").Find(&dbResult).Error; err == nil {
		for _, entity := range dbResult {
			if mdl, err := toModel(entity); err == nil {
				result = append(result, mdl)
			}
		}
	}
	return
}

func (repo *repository) CreateBookmark(model *model.Bookmark) (err error) {
	if entity, er := toEntity(model); er != nil {
		err = er
	} else {
		if err = repo.dbClient.Create(&entity).Error; err == nil {
			model.Id = entity.Id
		}
	}
	return
}

func (repo *repository) UpdateBookmark(model *model.Bookmark) (err error) {
	var entity entity.Bookmark
	if err = repo.dbClient.First(&entity, model.Id).Error; err == nil {
		if entity, err = toEntity(model); err == nil {
			return repo.dbClient.Save(entity).Error
		}
	} else {
		return &exception.NotFound{}
	}
	return
}

func (repo *repository) DeleteBookmark(id uuid.UUID) (err error) {
	return repo.dbClient.Delete(&entity.Bookmark{Id: &id}).Error
}
