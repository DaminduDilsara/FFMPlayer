package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func UploadFileApiController(c echo.Context) error {
	result, err := UploadFileApi(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}

func GetAllFilesApiController(c echo.Context) error {
	result, err := FindAllVideosApi(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	} else {
		return c.JSON(http.StatusOK, result)
	}
}
