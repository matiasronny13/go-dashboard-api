package bookmark

import (
	"io"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

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
	res, err := http.Get(info.DownloadUrl)
	if err != nil {
		slog.Error("error making http request: %s\n", err)
		return err
	}
	defer res.Body.Close()

	file, err := os.Create(filepath.Join(svc.appSettings.Api.FaviconFolder, info.FileName))
	if err != nil {
		slog.Error("Failed to create the file:", err)
		return err
	}
	defer file.Close()

	// Copy the image data from the HTTP response to the local file
	_, err = io.Copy(file, res.Body)
	if err != nil {
		slog.Error("Failed to save the image to the file:", err)
		return err
	}

	info.ImageUrl = "/logs/" + info.FileName
	return nil
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
