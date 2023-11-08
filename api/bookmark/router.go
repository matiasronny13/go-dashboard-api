package bookmark

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func BookmarkRouter(api *gin.RouterGroup, setting *model.AppSettings, dbClient *gorm.DB) {
	bookmarkHandler := Wire(setting, dbClient)
	bookmarkGroup := api.Group("/bookmark")
	{
		bookmarkGroup.POST("/favicon", bookmarkHandler.DownloadFavicon)
		bookmarkGroup.POST("", bookmarkHandler.CreateBookmark)
		bookmarkGroup.GET("", bookmarkHandler.GetBookmarks)
		bookmarkGroup.DELETE(":id", bookmarkHandler.DeleteBookmark)
		bookmarkGroup.PATCH(":id", bookmarkHandler.UpdateBookmark)
	}
}
