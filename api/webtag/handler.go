package webtag

import (
	"net/http"

	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/ghostrepo00/go-dashboard-api/util"
	"github.com/gin-gonic/gin"
)

type handler struct {
	appSettings *model.AppSettings
	service     WebtagService
}

type WebtagHandler interface {
	GetWebtag(c *gin.Context)
	GetWebtagById(c *gin.Context)
	CreateWebtag(c *gin.Context)
	UpdateWebtag(c *gin.Context)
	DeleteWebtag(c *gin.Context)
}

// GetWebtag godoc
// @Summary      Get all Webtag
// @Description  Get all Webtag
// @Tags         Webtag
// @Accept       json
// @Produce      json
// @Success      200  {array}  model.Webtag
// @Router       /webtag [get]
func (hdl *handler) GetWebtag(c *gin.Context) {
	if result, err := hdl.service.GetWebtag(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// GetWebtagById godoc
// @Summary      Get Webtag by id
// @Description  Get Webtag by id
// @Tags         Webtag
// @Accept       json
// @Produce      json
// @Param id path string true "Id"
// @Success      200  {object}  model.Webtag
// @Router       /webtag/{id} [get]
func (hdl *handler) GetWebtagById(c *gin.Context) {
	var (
		err    error
		result model.Webtag
	)

	if result, err = hdl.service.GetWebtagById(c.Param("id")); err == nil {
		c.JSON(http.StatusOK, result)
	}

	if err != nil {
		c.JSON(util.MapHttpError(err), gin.H{"error": err.Error()})
	}
}

// CreateWebtag godoc
// @Summary Create Webtag
// @Description Create Webtag
// @Tags Webtag
// @Accept application/json
// @Produce application/json
// @Param Webtag body model.Webtag true "New Webtag"
// @Success 201 {object} model.Webtag
// @Router /webtag [post]
func (hdl *handler) CreateWebtag(c *gin.Context) {
	var (
		Webtag model.Webtag
		err    error
	)

	if err = c.BindJSON(&Webtag); err == nil {
		if err = hdl.service.CreateWebtag(&Webtag); err == nil {
			c.JSON(http.StatusCreated, Webtag)
		}
	}

	if err != nil {
		c.JSON(util.MapHttpError(err), gin.H{"error": err.Error()})
	}
}

// UpdateWebtag godoc
// @Summary Update Webtag
// @Description Update Webtag
// @Tags Webtag
// @Accept application/json
// @Produce application/json
// @Param id path string true "Id"
// @Param Webtag body model.Webtag true "Existing Webtag"
// @Success 200 {object} model.Webtag
// @Router /webtag/{id} [patch]
func (hdl *handler) UpdateWebtag(c *gin.Context) {
	var (
		err error
		mdl model.Webtag
	)

	id := c.Param("id")
	if err = c.BindJSON(&mdl); err == nil {
		mdl.Id = id
		if err = hdl.service.UpdateWebtag(&mdl); err == nil {
			c.JSON(http.StatusOK, mdl)
		}
	}

	if err != nil {
		c.JSON(util.MapHttpError(err), gin.H{"error": err.Error()})
	}
}

// DeleteWebtag godoc
// @Summary Delete Webtag
// @Description Delete Webtag
// @Tags Webtag
// @Accept application/json
// @Produce application/json
// @Param id path string true "Id"
// @Success 200 {object} nil
// @Router /webtag/{id} [delete]
func (hdl *handler) DeleteWebtag(c *gin.Context) {
	var err error

	if err = hdl.service.DeleteWebtag(c.Param("id")); err == nil {
		c.JSON(http.StatusOK, nil)
	}

	if err != nil {
		c.JSON(util.MapHttpError(err), gin.H{"error": err.Error()})
	}
}
