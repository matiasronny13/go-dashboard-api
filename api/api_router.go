package api

import (
	swaggerDocs "github.com/ghostrepo00/go-dashboard-api/docs"
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/ghostrepo00/go-dashboard-api/feature/notes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
)

func SetupSwagger(router *gin.Engine, setting *model.AppSettings) {
	swaggerDocs.SwaggerInfo.Host = "localhost:8888"
	swaggerDocs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func SetupRouters(route *gin.Engine, setting *model.AppSettings, dbClient *gorm.DB) *gin.Engine {
	api := route.Group("/api/v1")
	notes.NotesRouter(api, setting, dbClient)
	return route
}
