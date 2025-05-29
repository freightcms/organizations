package web

import "github.com/freightcms/organizations/models"

type (
	GetAllOrganizationsRequest struct {
		// Limit the numer of results
		Limit int `json:"limit" query:"limit"`
		// Page number of the query
		Page int `json:"page" query:"page"`
	}

	// GetAllOrganizationsResponse is provided as the JSON or XML bindable repsonse
	// to an HTTP Request
	GetAllOrganizationsResponse struct {
		// Total is the number of results that are in the query
		Total         int64                  `json:"total" xml:"total"`
		Organizations []*models.Organization `json:"people" xml:"organizations"`
	}
)
