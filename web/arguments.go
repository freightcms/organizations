package web

import "github.com/graphql-go/graphql"

var (
	IDArgumentField = graphql.ArgumentConfig{
		Description: "Required argument for identifying the resource to be modified",
		Type:        graphql.NewNonNull(graphql.String),
	}
)
