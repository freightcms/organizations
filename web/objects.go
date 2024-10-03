package web

import (
	"github.com/graphql-go/graphql"
)

var (
	LocationObject *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "Location",
		Fields: graphql.Fields{
			"id":         &IDField,
			"line1":      &Line1Field,
			"line2":      &Line2Field,
			"line3":      &Line3Field,
			"locale":     &LocaleField,
			"region":     &RegionField,
			"country":    &CountryField,
			"postalCode": &PostalCodeField,
			"attention":  &AttentionField,
		},
	})
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
				Type: LocationObject,
			},
			"billingAddress": &graphql.Field{
				Name: "location",
				Type: LocationObject,
			},
		},
	})
)
