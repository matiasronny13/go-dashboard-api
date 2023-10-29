package notes

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/entity"
	"gorm.io/gorm"
)

type repository struct {
	dbClient *gorm.DB
}

type NotesRepository interface {
	GetNotes() (result []entity.Notes, err error)
}

func (repo *repository) GetNotes() (result []entity.Notes, err error) {
	return result, repo.dbClient.Find(&result).Error
}
