package notes

import (
	"net/http"

	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/ghostrepo00/go-dashboard-api/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	appSettings *model.AppSettings
	service     NotesService
}

type NotesHandler interface {
	GetNotes(c *gin.Context)
	GetNotesById(c *gin.Context)
	CreateNotes(c *gin.Context)
	UpdateNotes(c *gin.Context)
	DeleteNotes(c *gin.Context)
}

// GetNotes godoc
// @Summary      Get all notes
// @Description  Get all notes
// @Tags         Notes
// @Accept       json
// @Produce      json
// @Success      200  {array}  model.Notes
// @Router       /notes [get]
func (hdl *handler) GetNotes(c *gin.Context) {
	if result, err := hdl.service.GetNotes(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// GetNotesById godoc
// @Summary      Get notes by id
// @Description  Get notes by id
// @Tags         Notes
// @Accept       json
// @Produce      json
// @Param id path string true "Id" format(uuid)
// @Success      200  {object}  model.Notes
// @Router       /notes/{id} [get]
func (hdl *handler) GetNotesById(c *gin.Context) {
	var (
		err    error
		result model.Notes
		param  uuid.UUID
	)

	if param, err = uuid.Parse(c.Param("id")); err == nil {
		if result, err = hdl.service.GetNotesById(param); err == nil {
			c.JSON(http.StatusOK, result)
		}
	}

	if err != nil {
		c.JSON(util.MapHttpError(err), gin.H{"error": err.Error()})
	}
}

// CreateNotes godoc
// @Summary Create notes
// @Description Create notes
// @Tags Notes
// @Accept application/json
// @Produce application/json
// @Param notes body model.Notes true "New Notes"
// @Success 201 {object} model.Notes
// @Router /notes [post]
func (hdl *handler) CreateNotes(c *gin.Context) {
	var (
		notes model.Notes
		err   error
	)

	if err = c.BindJSON(&notes); err == nil {
		if err = hdl.service.CreateNotes(&notes); err == nil {
			c.JSON(http.StatusCreated, notes)
		}
	}

	if err != nil {
		c.JSON(util.MapHttpError(err), gin.H{"error": err.Error()})
	}
}

// UpdateNotes godoc
// @Summary Update notes
// @Description Update notes
// @Tags Notes
// @Accept application/json
// @Produce application/json
// @Param id path string true "Id" format(uuid)
// @Param notes body model.Notes true "Existing Notes"
// @Success 200 {object} model.Notes
// @Router /notes/{id} [patch]
func (hdl *handler) UpdateNotes(c *gin.Context) {
	var (
		id  uuid.UUID
		err error
		mdl model.Notes
	)

	param := c.Param("id")
	if id, err = uuid.Parse(param); err == nil {
		if err = c.BindJSON(&mdl); err == nil {
			mdl.Id = &id
			if err = hdl.service.UpdateNotes(&mdl); err == nil {
				c.JSON(http.StatusOK, mdl)
			}
		}
	}

	if err != nil {
		c.JSON(util.MapHttpError(err), gin.H{"error": err.Error()})
	}
}

// DeleteNotes godoc
// @Summary Delete notes
// @Description Delete notes
// @Tags Notes
// @Accept application/json
// @Produce application/json
// @Param id path string true "Id" format(uuid)
// @Success 200 {object} nil
// @Router /notes/{id} [delete]
func (hdl *handler) DeleteNotes(c *gin.Context) {
	var (
		err error
		id  uuid.UUID
	)

	param := c.Param("id")
	if id, err = uuid.Parse(param); err == nil {
		if err = hdl.service.DeleteNotes(id); err == nil {
			c.JSON(http.StatusOK, nil)
		}
	}

	if err != nil {
		c.JSON(util.MapHttpError(err), gin.H{"error": err.Error()})
	}
}
