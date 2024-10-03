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
		"line2":       "door #4",
		"line3":       "bin 5",
		"region":      "State, Province, or Territory",
		"locale":      "City, State or region",
		"postalCode":  "12345-2345",
		"countryCode": "US",
		"description": "My first little location",
	}
	// act
	addr := AddressFromArgs(locationModels.Billing, data)

	// assert
	assert.Equal(t, addr.Line1, data["line1"])
	assert.Equal(t, addr.Line2, data["line2"])
	assert.Equal(t, addr.Line3, data["line3"])
	assert.Equal(t, addr.Country, locationModels.CountryCode(data["countryCode"].(string)))
	assert.Equal(t, addr.Description, data["description"])
	assert.Equal(t, addr.Locale, data["locale"])
	assert.Equal(t, addr.PostalCode, data["postalCode"])
	assert.Equal(t, addr.Region, data["region"])
}
