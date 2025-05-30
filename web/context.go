package web

import (
	"github.com/freightcms/organizations/db"
	"github.com/labstack/echo/v4"
)

type (
	AppContext struct {
		echo.Context
		db.DbContext
	}
)
