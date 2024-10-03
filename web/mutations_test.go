package web

import (
	"testing"

	"gotest.tools/v3/assert"
)

func Test_AddressFromArgs_Should_SetFields(t *testing.T) {
	// arrange
	line2 := "door #4"
	line3 := "bin 5"

	data := map[string]interface{}{
		"line1":      "123 Fake Addr St. SW",
		"line2":      &line2,
		"line3":      &line3,
		"region":     "State, Province, or Territory",
		"locale":     "City, State or region",
		"postalCode": "12345-2345",
		"country":    "US",
	}
	// act
	addr := AddressFromArgs(data)

	// assert
	assert.Equal(t, addr.Line1, data["line1"])
	assert.Equal(t, addr.Line2, data["line2"])
	assert.Equal(t, addr.Line3, data["line3"])
	assert.Equal(t, addr.Country, data["country"])
	assert.Equal(t, addr.Description, data["description"])
	assert.Equal(t, addr.Local, data["locale"])
	assert.Equal(t, addr.PostalCode, data["postalCode"])
	assert.Equal(t, addr.Region, data["region"])

}
