package web

import (
	locationModels "github.com/freightcms/locations/models"
	locationWeb "github.com/freightcms/locations/web"
	"github.com/graphql-go/graphql"
	"github.com/freightcms/organizations/models"
)

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
		org.MailingAddress = locationWeb.AddressFromArgs(locationModels.Mailing, mailingAddress)
	}
	if _, ok := params.Args["billingAddress"]; ok {
		billingAddress := params.Args["billingAddress"].(map[string]interface{})
		org.BillingAddress = locationWeb.AddressFromArgs(locationModels.Billing, billingAddress)
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
