package web

import (
	"errors"
	"fmt"

	"github.com/graphql-go/graphql"

	locationModels "github.com/freightcms/locations/models"
	"github.com/squishedfox/organization-webservice/db/mongodb"
	"github.com/squishedfox/organization-webservice/models"
)

func AddressFromArgs(locationType locationModels.AddressType, args map[string]interface{}) *locationModels.AddressModel {
	model := &locationModels.AddressModel{
		Type:  locationType,
		Line1: args["line1"].(string),
	}
	if val, ok := args["line2"]; val != nil && ok {
		model.Line2 = val.(*string)
	}
	if val, ok := args["line3"]; val != nil && ok {
		model.Line3 = val.(*string)
	}
	if val, ok := args["locale"]; ok {
		model.Locale = val.(string)
	}
	if val, ok := args["countryCode"]; ok {
		model.Country = locationModels.CountryCode(val.(string))
	}
	if val, ok := args["region"]; ok {
		model.Region = val.(string)
	}
	if val, ok := args["postalCode"]; ok {
		model.PostalCode = val.(string)
	}
	if val, ok := args["attention"]; val != nil && ok {
		model.Attention = val.(*string)
	}
	if val, ok := args["description"]; val != nil && ok {
		model.Description = val.(*string)
	}
	return model
}

func OrganizationFromParams(params graphql.ResolveParams) *models.Organization {
	org := &models.Organization{}

	if _, ok := params.Args["name"]; ok {
		org.Name = params.Args["name"].(string)
	}
	if _, ok := params.Args["dba"]; ok {
		org.DBA = params.Args["dba"].(*string)
	}
	if _, ok := params.Args["rollupId"]; ok {
		org.RollupID = params.Args["rollupId"].(*string)
	}
	if _, ok := params.Args["mailingAddress"]; ok {
		mailingAddress := params.Args["mailingAddress"].(map[string]interface{})
		org.MailingAddress = AddressFromArgs(locationModels.Mailing, mailingAddress)
	}
	if _, ok := params.Args["billingAddress"]; ok {
		billingAddress := params.Args["billingAddress"].(map[string]interface{})
		org.BillingAddress = AddressFromArgs(locationModels.Billing, billingAddress)
	}
	return org
}

// MergeOrganization takes the first organization argument as the original organization
// and appends the params that were provided in the graphql query. If a parameter argument
// is provided for the field it is set.
func MergeOrganization(o *models.Organization, params graphql.ResolveParams) {
	if _, ok := params.Args["dba"]; ok {
		o.DBA = params.Args["dba"].(*string)
	}
	if _, ok := params.Args["name"]; ok {
		o.Name = params.Args["name"].(string)
	}
	if _, ok := params.Args["rollupId"]; ok {
		o.RollupID = params.Args["rollupId"].(*string)
	}
}

var (
	CreateLocationInput = graphql.NewInputObject(graphql.InputObjectConfig{
		Name: "CreateLocationInput",
		Fields: graphql.InputObjectConfigFieldMap{
			"line1": &graphql.InputObjectFieldConfig{
				Description: Line1Field.Description,
				Type:        Line1Field.Type,
			},
			"line2": &graphql.InputObjectFieldConfig{
				Description:  Line2Field.Description,
				Type:         Line2Field.Type,
				DefaultValue: nil,
			},
			"line3": &graphql.InputObjectFieldConfig{
				Description:  "Typically a bin within a warehouse or a floor number and building section",
				Type:         graphql.String,
				DefaultValue: nil,
			},
			"locale": &graphql.InputObjectFieldConfig{
				Description: LocaleField.Description,
				Type:        LocaleField.Type,
			},
			"region": &graphql.InputObjectFieldConfig{
				Description: RegionField.Description,
				Type:        RegionField.Type,
			},
			"postalCode": &graphql.InputObjectFieldConfig{
				Description: PostalCodeField.Description,
				Type:        PostalCodeField.Type,
			},
			"countryCode": &graphql.InputObjectFieldConfig{
				Description: CountryField.Description,
				Type:        CountryField.Type,
			},
			"attention": &graphql.InputObjectFieldConfig{
				Description: AttentionField.Description,
				Type:        AttentionField.Type,
			},
		},
	})
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
						Type:        CreateLocationInput,
					},
					"billingAddress": &graphql.ArgumentConfig{
						Description: "reliable address to forward invoices, shipping documents, or payments",
						Type:        CreateLocationInput,
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
