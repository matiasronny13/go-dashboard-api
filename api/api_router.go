package api

import (
	"github.com/ghostrepo00/go-dashboard-api/api/bookmark"
	"github.com/ghostrepo00/go-dashboard-api/api/notes"
	"github.com/ghostrepo00/go-dashboard-api/api/webtag"
	swaggerDocs "github.com/ghostrepo00/go-dashboard-api/docs"
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupSwaggerRouter(router *gin.Engine, setting *model.AppSettings) {
	swaggerDocs.SwaggerInfo.Host = setting.Api.Host
	swaggerDocs.SwaggerInfo.BasePath = setting.Api.BasePath
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func SetupApiRouters(router *gin.Engine, setting *model.AppSettings, dbClient *gorm.DB) {
	api := router.Group(setting.Api.BasePath)
	notes.NotesRouter(api, setting, dbClient)
	bookmark.BookmarkRouter(api, setting, dbClient)
	webtag.WebtagRouter(api, setting, dbClient)
}
