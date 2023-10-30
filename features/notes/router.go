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
		notesGroup.GET(":id", notesHandler.GetNotesById)
		notesGroup.POST("", notesHandler.CreateNotes)
		notesGroup.PATCH(":id", notesHandler.UpdateNotes)
		notesGroup.DELETE(":id", notesHandler.DeleteNotes)
	}
}
