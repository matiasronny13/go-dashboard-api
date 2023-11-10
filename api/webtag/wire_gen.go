// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package webtag

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"gorm.io/gorm"
)

// Injectors from wire.go:

func Wire(setting *model.AppSettings, dbClient *gorm.DB) *handler {
	webtagRepository := ProvideWebtagRepository(dbClient)
	webtagService := ProvideWebtagService(webtagRepository)
	webtagHandler := ProvideWebtagHandler(setting, webtagService)
	return webtagHandler
}