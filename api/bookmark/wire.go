//go:build wireinject
// +build wireinject

package bookmark

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func Wire(setting *model.AppSettings, dbClient *gorm.DB) *handler {
	wire.Build(wire.NewSet(
		ProvideBookmarkHandler,
		ProvideBookmarkService,
		ProvideBookmarkRepository,
		wire.Bind(new(BookmarkHandler), new(*handler)),
		wire.Bind(new(BookmarkService), new(*service)),
		wire.Bind(new(BookmarkRepository), new(*repository)),
	))
	return &handler{}
}
