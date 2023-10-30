package api

import (
	swaggerDocs "github.com/ghostrepo00/go-dashboard-api/docs"
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/ghostrepo00/go-dashboard-api/features/notes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupSwagger(router *gin.Engine, setting *model.AppSettings) {
	swaggerDocs.SwaggerInfo.Host = setting.Api.Host
	swaggerDocs.SwaggerInfo.BasePath = setting.Api.BasePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func SetupRouters(route *gin.Engine, setting *model.AppSettings, dbClient *gorm.DB) *gin.Engine {
	api := route.Group(setting.Api.BasePath)
	notes.NotesRouter(api, setting, dbClient)
	return route
}
