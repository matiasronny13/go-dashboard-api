package bookmark

import (
	"errors"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ghostrepo00/go-dashboard-api/domain/model"
	"github.com/google/uuid"
)

type service struct {
	appSettings *model.AppSettings
	repository  BookmarkRepository
}

type BookmarkService interface {
	DownloadFavicon(info *model.Favicon) (err error)
	GetBookmarks() (result []model.Bookmark, err error)
	CreateBookmark(model *model.Bookmark) (err error)
	UpdateBookmark(model *model.Bookmark) (err error)
	DeleteBookmark(id uuid.UUID) (err error)
}

func (svc *service) DownloadFavicon(info *model.Favicon) (err error) {
	var resp *http.Response

	if resp, err = http.Get(info.DownloadUrl); err == nil {
		defer resp.Body.Close()
		if strings.HasPrefix(strings.ToLower(resp.Header.Get("Content-Type")), "image/") {
			if file, err := os.Create(filepath.Join(svc.appSettings.Api.FaviconFolder, info.FileName)); err == nil {
				defer file.Close()

				// Copy the image data from the HTTP response to the local file
				if _, err = io.Copy(file, resp.Body); err == nil {
					info.ImageUrl = "/logs/" + info.FileName
				}
			}
		} else {
			return errors.New("file type is not image")
		}
	}
	return
}

func (svc *service) GetBookmarks() (result []model.Bookmark, err error) {
	return svc.repository.GetBookmarks()
}

func (svc *service) CreateBookmark(model *model.Bookmark) (err error) {
	return svc.repository.CreateBookmark(model)
}

func (svc *service) DeleteBookmark(id uuid.UUID) (err error) {
	return svc.repository.DeleteBookmark(id)
}

func (svc *service) UpdateBookmark(model *model.Bookmark) (err error) {
	return svc.repository.UpdateBookmark(model)
}
