package web

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"
	"github.com/squishedfox/webservice-prototype/db/mongodb"
	"github.com/squishedfox/webservice-prototype/models"
)

func OrganizationFromParams(params graphql.ResolveParams) *models.Organization {
	return &models.Organization{
		ID:       "",
		Name:     params.Args["name"].(string),
		DBA:      params.Args["dba"].(string),
		RollupID: params.Args["rollupId"].(string),
	}
}

func MergeOrganization(o *models.Organization, params graphql.ResolveParams) {
	if _, ok := params.Args["dba"]; ok {
		o.DBA = params.Args["dba"].(string)
	}
	if _, ok := params.Args["name"]; ok {
		o.Name = params.Args["name"].(string)
	}
	if _, ok := params.Args["rollupId"]; ok {
		o.RollupID = params.Args["rollupId"].(string)
	}
}

var (
	Mutations *graphql.Object = graphql.NewObject(graphql.ObjectConfig{
		Name: "mutations",
		Fields: graphql.Fields{
			"createOrganization": &graphql.Field{
				Type:        IDObject,
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
						Description: RollupID.Description,
						Type:        RollupID.Type,
					},
				},
				Resolve: func(params graphql.ResolveParams) (interface{}, error) {
					model := OrganizationFromParams(params)
					mgr := mongodb.FromContext(params.Context)
					id, err := mgr.CreateOrganization(model)
					if err != nil {
						return nil, err
					}
					return id, err
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
						Type: graphql.String,
					},
					"name": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
					"rollupId": &graphql.ArgumentConfig{
						Type: graphql.String,
					},
				}, // ends aarguments
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
