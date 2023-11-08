package notes

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/entity"
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/ghostrepo00/go-dashboard-api/domain/model/exception"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	dbClient *gorm.DB
}

type NotesRepository interface {
	GetNotes() (result []model.Notes, err error)
	GetNotesById(id uuid.UUID) (result model.Notes, err error)
	CreateNotes(model *model.Notes) (err error)
	UpdateNotes(model *model.Notes) (err error)
	DeleteNotes(id uuid.UUID) (err error)
}

func (repo *repository) GetNotes() (result []model.Notes, err error) {
	var dbResult []entity.Notes
	if err = repo.dbClient.Find(&dbResult).Error; err == nil {
		for _, v := range dbResult {
			result = append(result, model.Notes{Id: v.Id, Title: v.Title, Content: v.Content})
		}
	}
	return
}

func (repo *repository) GetNotesById(id uuid.UUID) (result model.Notes, err error) {
	var dbResult *entity.Notes
	if err = repo.dbClient.First(&dbResult, id).Error; err == nil {
		result = model.Notes{Id: dbResult.Id, Title: dbResult.Title, Content: dbResult.Content}
	} else {
		return model.Notes{}, &exception.NotFound{}
	}
	return
}

func (repo *repository) CreateNotes(model *model.Notes) (err error) {
	entity := entity.Notes{Title: model.Title, Content: model.Content}
	if err = repo.dbClient.Create(&entity).Error; err == nil {
		model.Id = entity.Id
	}
	return
}

func (repo *repository) UpdateNotes(model *model.Notes) (err error) {
	var entity entity.Notes
	if err = repo.dbClient.First(&entity, model.Id).Error; err == nil {
		entity.Title = model.Title
		entity.Content = model.Content
		return repo.dbClient.Save(entity).Error
	} else {
		return &exception.NotFound{}
	}
}

func (repo *repository) DeleteNotes(id uuid.UUID) (err error) {
	return repo.dbClient.Delete(&entity.Notes{Id: &id}).Error
}
