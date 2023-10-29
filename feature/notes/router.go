package notes

import (
	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NotesRouter(api *gin.RouterGroup, setting *model.AppSettings, dbClient *gorm.DB) {
	notesHandler := Wire(setting, dbClient)
	notesGroup := api.Group("/notes")
	{
		notesGroup.GET("", notesHandler.GetNotes)
		// notesGroup.GET(":id", c.ShowAccount)
		// notesGroup.POST("", c.AddAccount)
		// notesGroup.POST(":id/images", c.UploadAccountImage)
		// notesGroup.DELETE(":id", c.DeleteAccount)
		// notesGroup.PATCH(":id", c.UpdateAccount)
	}
}
