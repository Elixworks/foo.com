package controllers

import (
	"net/http"

	"format-foo.com/models"
	"github.com/labstack/echo"
)

func GetInfo(c echo.Context) error {
	var s models.SiteInfo
	s.New()

	return c.JSON(http.StatusOK, H{
		"shopInfo": s,
	})
}
