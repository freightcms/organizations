package web

import (
	"net/http"

	"github.com/freightcms/organizations/db"
	"github.com/labstack/echo/v4"
)

func getAllOrganizationsHandler(c echo.Context) error {
	var r GetAllOrganizationsRequest
	if err := c.Bind(&r); err != nil {
		return err
	}
	q := db.NewQuery().SetPage(r.Page).SetPageSize(r.Limit)
	organizations, count, err := c.(AppContext).OrganizationResourceManager.Get(q)
	if err != nil {
		return err
	}
	res := GetAllOrganizationsResponse{
		Total:         count,
		Organizations: organizations,
	}
	c.JSON(http.StatusOK, res)
	return nil
}
