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

func Test_AddressFromArgs_Should_SetOptionalFields(t *testing.T) {
	// arrange
	var line2 string = "a suite or office"
	var line3 string = "a floor, bin, warehouse section, etc"
	var notes string = "please validate fields"
	var attention string = "unit testing person"
	var description string = "A unit test for optioal fields"
	data := map[string]interface{}{
		"line1":       "123 Fake Addr St. SW",
		"line2":       &line2,
		"line3":       &line3,
		"region":      "State, Province, or Territory",
		"locale":      "City, State or region",
		"postalCode":  "12345-2345",
		"countryCode": "US",
		"description": &description,
		"attention":   &attention,
		"notes":       &notes,
	}
	// act
	addr := AddressFromArgs(locationModels.Billing, data)

	// assert
	assert.Equal(t, addr.Line1, data["line1"].(string))
	assert.Equal(t, addr.Line2, data["line2"].(*string))
	assert.Equal(t, addr.Line3, data["line3"].(*string))
	assert.Equal(t, addr.Country, locationModels.CountryCode(data["countryCode"].(string)))
	assert.Equal(t, addr.Description, data["description"].(*string))
	assert.Equal(t, addr.Locale, data["locale"].(string))
	assert.Equal(t, addr.PostalCode, data["postalCode"].(string))
	assert.Equal(t, addr.Region, data["region"].(string))
	assert.Equal(t, addr.Attention, data["attention"].(*string))
	assert.Equal(t, addr.Type, locationModels.Billing)
	assert.Equal(t, addr.Notes, data["notes"].(*string))
}
