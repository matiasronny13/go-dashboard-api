//go:build wireinject
// +build wireinject

package notes

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func Wire(setting *model.AppSettings, dbClient *gorm.DB) *handler {
	wire.Build(wire.NewSet(
		ProvideNotesHandler,
		ProvideNotesService,
		ProvideNotesRepository,
		wire.Bind(new(NotesHandler), new(*handler)),
		wire.Bind(new(NotesService), new(*service)),
		wire.Bind(new(NotesRepository), new(*repository)),
	))
	return &handler{}
}
