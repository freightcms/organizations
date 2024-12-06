package web

import (
	"errors"

	"github.com/graphql-go/graphql"
	"organization/db"
	"organizationdb/mongodb"
)

var (
	RootQuery *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "RootQuery",
		Fields: graphql.Fields{
			"organization": &graphql.Field{
				Type: OrganizationObject,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type:        IDField.Type,
						Description: "identifier of organization to find. If this value is provided other arguments such as page, pageSize, and sort by are ignored",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					mgr := mongodb.FromContext(p.Context)
					if _, ok := p.Args["id"]; ok {
						person, err := mgr.GetById(p.Args["id"].(string))
						return person, err
					}
					return nil, errors.New("missing filter")
				},
			},
			"organizations": &graphql.Field{
				Type: graphql.NewList(OrganizationObject),
				Args: graphql.FieldConfigArgument{
					"page": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 0,
					},
					"pageSize": &graphql.ArgumentConfig{
						Type:         graphql.Int,
						DefaultValue: 10,
					},
					"sortBy": &graphql.ArgumentConfig{
						Type: graphql.NewEnum(graphql.EnumConfig{
							Name:        "SortByFields",
							Description: "Valid field to sort data by in the return",
							Values: graphql.EnumValueConfigMap{
								"id": &graphql.EnumValueConfig{
									Value: "_id",
								},
							},
						}),
						DefaultValue: "_id",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					mgr := mongodb.FromContext(p.Context)
					q := db.NewQuery()
					if _, ok := p.Args["page"]; ok {
						q.SetPage(p.Args["page"].(int))
					}
					if _, ok := p.Args["pageSize"]; ok {
						q.SetPageSize(p.Args["pageSize"].(int))
					}
					if _, ok := p.Args["sortBy"]; ok {
						q.SetSortBy(p.Args["sortBy"].(string))
					}
					people, err := mgr.Get(q)
					return people, err
				},
			}, // end people field
		}, // end Fields
	}) // ends object
) // end var
