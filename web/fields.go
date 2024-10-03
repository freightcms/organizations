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
	RollupIDField = graphql.Field{
		Name:        "rollupId",
		Description: "Identifier of the parent/umbrella and legal Organization/Entity",
		Type:        graphql.String,
	}
	Line1Field = graphql.Field{
		Name:        "line1",
		Description: "typically the street address of a location",
		Type:        graphql.NewNonNull(graphql.String),
	}
	Line2Field = graphql.Field{
		Name:        "line2",
		Description: "typically a floor, suite, apartment, etc.",
		Type:        graphql.String,
	}
	Line3Field = graphql.Field{
		Name:        "line3",
		Description: "typically a bin or storage number in a warehouse. Sometimes this will also include the floor",
		Type:        graphql.String,
	}
	LocaleField = graphql.Field{
		Name:        "local",
		Description: "e.g. city, town, village, etc.",
		Type:        graphql.NewNonNull(graphql.String),
	}
	RegionField = graphql.Field{
		Name:        "region",
		Description: "state, province, or area within a country",
		Type:        graphql.NewNonNull(graphql.String),
	}
	PostalCodeField = graphql.Field{
		Name:        "postalCode",
		Description: "typically a five digit identifier followed by five or four specific mailing local",
		Type:        graphql.NewNonNull(graphql.String),
	}
	CountryField = graphql.Field{
		Name:        "country",
		Description: "two digit code ISO 3166-1 to identify country",
		Type:        graphql.NewNonNull(graphql.String),
	}
	AttentionField = graphql.Field{
		Name:        "attention",
		Description: "person or business to address item(s) to",
		Type:        graphql.String,
	}
)
