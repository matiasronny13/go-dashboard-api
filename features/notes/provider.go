package notes

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"gorm.io/gorm"
)

func ProvideNotesRepository(dbClient *gorm.DB) *repository {
	return &repository{dbClient: dbClient}
}

func ProvideNotesService(repository NotesRepository) *service {
	return &service{repository}
}

func ProvideNotesHandler(setting *model.AppSettings, service NotesService) *handler {
	return &handler{setting, service}
}
