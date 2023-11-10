package webtag

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func WebtagRouter(api *gin.RouterGroup, setting *model.AppSettings, dbClient *gorm.DB) {
	WebtagHandler := Wire(setting, dbClient)
	WebtagGroup := api.Group("/webtag")
	{
		WebtagGroup.GET("", WebtagHandler.GetWebtag)
		WebtagGroup.GET(":id", WebtagHandler.GetWebtagById)
		WebtagGroup.POST("", WebtagHandler.CreateWebtag)
		WebtagGroup.PATCH(":id", WebtagHandler.UpdateWebtag)
		WebtagGroup.DELETE(":id", WebtagHandler.DeleteWebtag)
	}
}
