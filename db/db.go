package db

import "github.com/squishedfox/organization-webservice/models"

type OrganizationQuery struct {
	// Page to start sorting by. Indexing at 1
	Page int
	// PageSize tells the query how many results to return based on search criteria
	PageSize int
	// SortBy Tells the query how it should be sorting the results
	SortBy string
	// Fields to include in the return statement
	Fields []string
}

// NewQuery creates a new query object the default values. This is the preferred method for creating
// a new query object. You should then set the other values of the query struct as necessary.
func NewQuery() *OrganizationQuery {
	return &OrganizationQuery{
		Page:     0,
		PageSize: 10,
		SortBy:   "_id",
		Fields:   []string{},
	}
}

func (q *OrganizationQuery) SetPage(page int) *OrganizationQuery {
	q.Page = page
	return q
}

func (q *OrganizationQuery) SetPageSize(pageSize int) *OrganizationQuery {
	q.PageSize = pageSize
	return q
}

func (q *OrganizationQuery) SetSortBy(sortBy string) *OrganizationQuery {
	q.SortBy = sortBy
	return q
}

func (q *OrganizationQuery) SetFields(fields []string) *OrganizationQuery {
	q.Fields = fields
	return q
}

// OrganizationResourceManager provides an abstract interface for managing Organization Resources to a database provider such as
// postgres, ms sql server, couchdb, mongodb, redis, dynamodb, etc.
type OrganizationResourceManager interface {
	// CreateOrganization function puts a new Organization resource into the database and returns the ID of the newly
	// created resource. if there is an error while attempting to create the Person resource it is
	// returned with a nil for the ID.
	CreateOrganization(model *models.Organization) (interface{}, error)

	// DeleteOrganization deletes an Organization resource from the target database system. If there is an error attempting
	// to delete the resource the error is returned. If the resource does not exist no error is returned.
	DeleteOrganization(id interface{}) error

	// UpdateOrganization modifies and updates an Organization resource. If there is an error attempting to update the
	// resource or a resource could not be found an error is returned.
	UpdateOrganization(id interface{}, model *models.Organization) error

	// GetById fetches an Organization resource by it's identifier. If no resource is found then nil, nil is returned
	// as a successfully "failed" attempt. If there is an issue communicating with the database system the error
	// is returned and nil for the resource.
	GetById(id interface{}) (*models.Organization, error)

	// Get fetches all Organization resources from target database/resource storage. If none are found an empty slice
	// is returned. If there is an error fetching one or more recrods the error is immediately returned at the
	// opperation is cancelled.
	Get(query *OrganizationQuery) ([]*models.Organization, error)

	// TODO: add query availability as well so we can search for resources based on properties
}
