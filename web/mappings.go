package web

import (
	locationModels "github.com/freightcms/locations/models"
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
	if val, ok := args["notes"]; val != nil && ok {
		model.Notes = val.(*string)
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


