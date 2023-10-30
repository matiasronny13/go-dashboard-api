package notes

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/entity"
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type repository struct {
	dbClient *gorm.DB
}

type NotesRepository interface {
	GetNotes() (result []model.Notes, err error)
	GetNotesById(id uuid.UUID) (result model.Notes, err error)
	CreateNotes(model *model.Notes) error
	UpdateNotes(model *model.Notes) error
	DeleteNotes(id uuid.UUID) error
}

func (repo *repository) GetNotes() (result []model.Notes, err error) {
	var dbResult []entity.Notes
	err = repo.dbClient.Find(&dbResult).Error
	for _, v := range dbResult {
		result = append(result, model.Notes{Id: v.Id, Title: v.Title, Content: v.Content})
	}
	return
}

func (repo *repository) GetNotesById(id uuid.UUID) (result model.Notes, err error) {
	var dbResult *entity.Notes
	err = repo.dbClient.Find(&dbResult).Error
	if dbResult != nil {
		result = model.Notes{Id: dbResult.Id, Title: dbResult.Title, Content: dbResult.Content}
	}
	return
}

func (repo *repository) CreateNotes(model *model.Notes) error {
	entity := entity.Notes{Title: model.Title, Content: model.Content}
	if err := repo.dbClient.Create(&entity).Error; err != nil {
		return err
	}
	model.Id = entity.Id
	return nil
}

func (repo *repository) UpdateNotes(model *model.Notes) error {
	var entity entity.Notes
	err := repo.dbClient.Table("notes").First(&entity, model.Id).Error
	if err != nil {
		return err
	}
	entity.Title = model.Title
	entity.Content = model.Content
	if err := repo.dbClient.Table("notes").Save(entity).Error; err != nil {
		return err
	}
	return nil
}

func (repo *repository) DeleteNotes(id uuid.UUID) error {
	if err := repo.dbClient.Table("notes").Where("id = ?", id).Delete(&entity.Notes{}).Error; err != nil {
		return err
	}
	return nil
}
