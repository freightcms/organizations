package web

import (
	locationWeb "github.com/freightcms/locations/web"
	"github.com/graphql-go/graphql"
)

var (
	IDObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "ID",
		Fields: graphql.Fields{
			"id": &IDField,
		},
	})
	OrganizationObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "Organization",
		Fields: graphql.Fields{
			"id":       &IDField,
			"dba":      &DBAField,
			"name":     &NameField,
			"rollupId": &RollupIDField,
			"mailingAddress": &graphql.Field{
				Name: "location",
				Type: locationWeb.LocationObject,
			},
			"billingAddress": &graphql.Field{
				Name: "location",
				Type: locationWeb.LocationObject,
			},
		},
	})
)
