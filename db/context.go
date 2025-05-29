package db

import "context"

type (
	DbContext struct {
		context.Context
		OrganizationResourceManager
	}
)
