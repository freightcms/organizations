package web

import "github.com/graphql-go/graphql"

var (
	IDField = graphql.Field{
		Name:        "id",
		Description: "Unique Identifier to access resource",
		Type:        graphql.NewNonNull(graphql.String),
	}
	DBAField = graphql.Field{
		Name:        "dba",
		Description: "Doing Business As name for an Organization. At times this will be blank or null when the Organization is using the same name",
		Type:        graphql.String,
	}
	NameField = graphql.Field{
		Name:        "name",
		Description: "Official (Legal) Name for an Organization/Entity",
		Type:        graphql.NewNonNull(graphql.String),
	}
	RollupID = graphql.Field{
		Name:        "rollupId",
		Description: "Identifier of the parent/umbrella and legal Organization/Entity",
		Type:        graphql.String,
	}
)
