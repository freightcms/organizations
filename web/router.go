package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Register(e *echo.Echo) {
	e.GET("/organizations", getAllOrganizationsHandler)
	e.GET("/status", echo.HandlerFunc(func(c echo.Context) error {
		body := struct {
			Status string `json:"status" xml:"status"`
		}{
			Status: "Ok",
		}
		return c.JSONPretty(http.StatusOK, &body, "    ")
	}))
}
