package util

import (
	"net/http"

	"github.com/ghostrepo00/go-dashboard-api/domain/model/exception"
)

func MapHttpError(err error) int {
	httpCode := http.StatusInternalServerError
	if _, ok := err.(*exception.NotFound); ok {
		httpCode = http.StatusNotFound
	}

	return httpCode
}
