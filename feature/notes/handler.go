package notes

import (
	"net/http"

	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/gin-gonic/gin"
)

type handler struct {
	appSettings *model.AppSettings
	service     NotesService
}

type NotesHandler interface {
	GetNotes(c *gin.Context)
}

// GetNotes godoc
// @Summary      Get all notes
// @Description  Get all notes
// @Tags         notes
// @Accept       json
// @Produce      json
// @Success      200  {object}  entity.Notes
// @Router       /notes [get]
func (hdl *handler) GetNotes(c *gin.Context) {
	result, err := hdl.service.GetNotes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, result)
	}
}
