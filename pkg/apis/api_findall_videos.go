package api

import (
	"AdvancedNetwork/pkg/dbOps"
	"AdvancedNetwork/pkg/models"
	"github.com/labstack/echo/v4"
)

func FindAllVideosApi(c echo.Context) ([]models.Video, error) {

	returnValue, err := dbOps.DB_GetAll_Videos()

	if err != nil {
		return returnValue, err
	}
	return returnValue, err
}
