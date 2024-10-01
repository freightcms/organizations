package web

import (
	"github.com/graphql-go/graphql"
)

var (
	LocationObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"id": IDField,
			"line1": &graphql.Field{
				Name:        "line1",
				Description: "typically the street address of a location",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"line2": &graphql.Field{
				Name:        "line2",
				Description: "typically a floor, suite, apartment, etc.",
				Type:        graphql.String,
			},
			"line3": &graphql.Field{
				Name:        "line3",
				Description: "typically a bin or storage number in a warehouse. Sometimes this will also include the floor",
				Type:        graphql.String,
			},
			"local": &graphql.Field{
				Name:        "local",
				Description: "e.g. city, town, village, etc.",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"region": &graphql.Field{
				Name:        "region",
				Description: "state, province, or area within a country",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"postalCode": &graphql.Field{
				Name:        "postalCode",
				Description: "typically a five digit identifier followed by five or four specific mailing local",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"country": &graphql.Field{
				Name:        "country",
				Description: "two digit code ISO 3166-1 to identify country",
				Type:        graphql.NewNonNull(graphql.String),
			},
			"attention": &graphql.Field{
				Name:        "attention",
				Description: "person or business to address item(s) to",
				Type:        graphql.String,
			},
		},
	})
	IDObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "ID",
		Fields: graphql.Fields{
			"id": &graphql.Field{
				Type: graphql.String,
			},
		},
	})
	OrganizationObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "Organization",
		Fields: graphql.Fields{
			"id":       IDField,
			"dba":      DBAField,
			"name":     NameField,
			"rollupId": RollupID,
			"mailingAddress": &graphql.Field{
				Name: "location",
				Type: LocationObject,
			},
			"billingAddress": &graphql.Field{
				Name: "location",
				Type: LocationObject,
			},
		},
	})
)
