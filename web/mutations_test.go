package web

import (
	locationModels "github.com/freightcms/locations/models"
	"gotest.tools/v3/assert"
	"testing"
)

func Test_AddressFromArgs_Should_SetFields(t *testing.T) {
	// arrange
	data := map[string]interface{}{
		"line1":       "123 Fake Addr St. SW",
		"line2":       nil,
		"line3":       nil,
		"region":      "State, Province, or Territory",
		"locale":      "City, State or region",
		"postalCode":  "12345-2345",
		"countryCode": "US",
		"description": nil,
		"attention":   nil,
		"notes":       nil,
	}
	// act
	addr := AddressFromArgs(locationModels.Billing, data)

	// assert
	assert.Equal(t, addr.Line1, data["line1"].(string))
	assert.Assert(t, addr.Line2 == nil)
	assert.Assert(t, addr.Line3 == nil)
	assert.Equal(t, addr.Country, locationModels.CountryCode(data["countryCode"].(string)))
	assert.Assert(t, addr.Description == nil)
	assert.Equal(t, addr.Locale, data["locale"].(string))
	assert.Equal(t, addr.PostalCode, data["postalCode"].(string))
	assert.Equal(t, addr.Region, data["region"].(string))
	assert.Assert(t, addr.Attention == nil)
	assert.Equal(t, addr.Type, locationModels.Billing)
	assert.Assert(t, addr.Notes == nil)
}
