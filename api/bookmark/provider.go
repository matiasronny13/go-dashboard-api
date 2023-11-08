package bookmark

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"gorm.io/gorm"
)

func ProvideBookmarkRepository(dbClient *gorm.DB) *repository {
	return &repository{dbClient: dbClient}
}

func ProvideBookmarkService(setting *model.AppSettings, repository BookmarkRepository) *service {
	return &service{setting, repository}
}

func ProvideBookmarkHandler(setting *model.AppSettings, service BookmarkService) *handler {
	return &handler{setting, service}
}
