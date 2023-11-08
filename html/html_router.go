package html

import (
	"net/http"

	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupHtmlRouters(router *gin.Engine, setting *model.AppSettings, dbClient *gorm.DB) {
	router.LoadHTMLGlob("html/templates/*")

	router.Static("/assets", "html/assets")

	router.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{"title": "Home"})
	})

	router.GET("/template/test", func(c *gin.Context) {
		c.HTML(http.StatusOK, "test.html", gin.H{"value": "value 1", "value2": "value 2"})
	})
}
