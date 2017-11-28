package controllers

import (
	"net/http"

	"format-foo.com/models"
	"github.com/labstack/echo"
)

func Index(c echo.Context) error {
	var s models.SiteInfo
	s.New()

	return c.Render(http.StatusOK, "primary-format.gohtml", H{
		"shopInfo": s,
	})
}
