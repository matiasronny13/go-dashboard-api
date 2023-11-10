package webtag

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"gorm.io/gorm"
)

func ProvideWebtagRepository(dbClient *gorm.DB) *repository {
	return &repository{dbClient: dbClient}
}

func ProvideWebtagService(repository WebtagRepository) *service {
	return &service{repository}
}

func ProvideWebtagHandler(setting *model.AppSettings, service WebtagService) *handler {
	return &handler{setting, service}
}
