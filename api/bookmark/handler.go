package bookmark

import (
	"net/http"

	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/ghostrepo00/go-dashboard-api/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type handler struct {
	appSettings *model.AppSettings
	service     BookmarkService
}

type BookmarkHandler interface {
	DownloadFavicon(c *gin.Context)
	GetBookmarks(c *gin.Context)
	CreateBookmark(c *gin.Context)
	UpdateBookmark(c *gin.Context)
	DeleteBookmark(c *gin.Context)
}

// DownloadFavicon godoc
// @Summary Download favicon to server folder
// @Description Download favicon to server folder
// @Tags Bookmark
// @Accept application/json
// @Produce application/json
// @Param info body model.Favicon true "Favicon info"
// @Success 200 {object} model.Favicon
// @Router /bookmark/favicon [post]
func (hdl *handler) DownloadFavicon(c *gin.Context) {
	var info model.Favicon
	if err := c.BindJSON(&info); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		if err := hdl.service.DownloadFavicon(&info); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusCreated, info)
		}
	}
}

// GetBookmarks godoc
// @Summary      Get all bookmarks
// @Description  Get all bookmarks
// @Tags         Bookmark
// @Accept       json
// @Produce      json
// @Success      200  {array}  model.Bookmark
// @Router       /bookmark [get]
func (hdl *handler) GetBookmarks(c *gin.Context) {
	if result, err := hdl.service.GetBookmarks(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// CreateBookmark godoc
// @Summary Create Bookmark
// @Description Create Bookmark
// @Tags Bookmark
// @Accept application/json
// @Produce application/json
// @Param info body model.Bookmark true "Bookmark info"
// @Success 201 {object} model.Bookmark
// @Router /bookmark [post]
func (hdl *handler) CreateBookmark(c *gin.Context) {
	var (
		model model.Bookmark
		err   error
	)

	if err := c.BindJSON(&model); err == nil {
		if err := hdl.service.CreateBookmark(&model); err == nil {
			c.JSON(http.StatusCreated, model)
		}
	}

	if err != nil {
		c.JSON(util.MapHttpError(err), gin.H{"error": err.Error()})
	}
}

// UpdateBookmark godoc
// @Summary Update Bookmark
// @Description Update Bookmark
// @Tags Bookmark
// @Accept application/json
// @Produce application/json
// @param id path string true "Id" format(uuid)
// @Param info body model.Bookmark true "Bookmark info"
// @Success 200 {object} model.Bookmark
// @Router /bookmark/{id} [patch]
func (hdl *handler) UpdateBookmark(c *gin.Context) {
	var (
		id       uuid.UUID
		err      error
		bookmark model.Bookmark
	)

	param := c.Param("id")
	if id, err = uuid.Parse(param); err == nil {
		if err = c.BindJSON(&bookmark); err == nil {
			bookmark.Id = &id
			if err = hdl.service.UpdateBookmark(&bookmark); err == nil {
				c.JSON(http.StatusOK, bookmark)
			}
		}
	}

	if err != nil {
		c.JSON(util.MapHttpError(err), gin.H{"error": err.Error()})
	}
}

// DeleteBookmark godoc
// @Summary Delete Bookmark
// @Description Delete Bookmark
// @Tags Bookmark
// @Accept application/json
// @Produce application/json
// @Param id path string true "Id" format(uuid)
// @Success 200 {object} nil
// @Router /bookmark/{id} [delete]
func (hdl *handler) DeleteBookmark(c *gin.Context) {
	var (
		err error
		id  uuid.UUID
	)

	param := c.Param("id")
	if id, err = uuid.Parse(param); err == nil {
		if err = hdl.service.DeleteBookmark(id); err == nil {
			c.JSON(http.StatusOK, nil)
		}
	}

	if err != nil {
		c.JSON(util.MapHttpError(err), gin.H{"error": err.Error()})
	}
}
