//go:build wireinject
// +build wireinject

package webtag

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func Wire(setting *model.AppSettings, dbClient *gorm.DB) *handler {
	wire.Build(wire.NewSet(
		ProvideWebtagHandler,
		ProvideWebtagService,
		ProvideWebtagRepository,
		wire.Bind(new(WebtagHandler), new(*handler)),
		wire.Bind(new(WebtagService), new(*service)),
		wire.Bind(new(WebtagRepository), new(*repository)),
	))
	return &handler{}
}
