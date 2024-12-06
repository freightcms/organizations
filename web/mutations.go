package web

import (
	"errors"
	"fmt"

	locationWeb "github.com/freightcms/locations/web"
	"github.com/graphql-go/graphql"
	"organizations/db/mongodb"
)

var (
	Mutations = graphql.NewObject(graphql.ObjectConfig{
		Name: "mutations",
		Fields: graphql.Fields{
			"createOrganization": &graphql.Field{
				Name: "CreateOrganization",
				Type: graphql.NewObject(graphql.ObjectConfig{
					Name:        "IDResponse",
					Description: "receive this response during a simple operation",
					Fields: graphql.Fields{
						"id": &IDField,
					},
				}),
				Description: "Create new Organization",
				Args: graphql.FieldConfigArgument{
					"dba": &graphql.ArgumentConfig{
						Description: DBAField.Description,
						Type:        DBAField.Type,
					},
					"name": &graphql.ArgumentConfig{
						Description: NameField.Description,
						Type:        NameField.Type,
					},
					"rollupId": &graphql.ArgumentConfig{
						Description: RollupIDField.Description,
						Type:        RollupIDField.Type,
					},
					"mailingAddress": &graphql.ArgumentConfig{
						Description: "Can be any address which should be a physical location",
						Type:        locationWeb.CreateLocationInput,
					},
					"billingAddress": &graphql.ArgumentConfig{
						Description: "reliable address to forward invoices, shipping documents, or payments",
						Type:        locationWeb.CreateLocationInput,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					model := OrganizationFromParams(params)
					mgr := mongodb.FromContext(params.Context)
					id, err := mgr.CreateOrganization(model)
					if err != nil {
						return nil, err
					}
					resp := struct {
						ID string `json:"id" bson:"id"`
					}{
						ID: id.(string),
					}
					return &resp, err
				},
			},
			"deleteOrganization": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "Delete an existing Person resource",
				Args: graphql.FieldConfigArgument{
					"id": &IDArgumentField,
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					mgr := mongodb.FromContext(params.Context)
					if _, ok := params.Args["id"]; !ok {
						return false, errors.New("no Resource with the ID exists")
					}
					err := mgr.DeleteOrganization(params.Args["id"].(string))
					return true, err
				},
			},
			"UpdateOrganization": &graphql.Field{
				Type:        graphql.Boolean,
				Description: "Update an existing Organization object. All fields exceept for the ID field are optional since they are not set if they are not provided in the query.",
				Args: graphql.FieldConfigArgument{
					"id": &IDArgumentField,
					"dba": &graphql.ArgumentConfig{
						Description: DBAField.Description,
						Type:        graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Description: NameField.Description,
						Type:        graphql.String,
					},
					"rollupId": &graphql.ArgumentConfig{
						Description: RollupIDField.Description,
						Type:        graphql.String,
					},
				}, // ends arguments
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					mgr := mongodb.FromContext(params.Context)
					id := params.Args["id"].(string)
					model, err := mgr.GetById(id)
					if err != nil {
						return nil, err
					}
					if model == nil {
						return nil, fmt.Errorf("could not find person with ID %s", id)
					}
					MergeOrganization(model, params)

					if err := mgr.UpdateOrganization(id, model); err != nil {
						return nil, err
					}
					return true, nil
				}, // end Resolve
			}, // ends updatePerson Field type definition
		},
	})
)
