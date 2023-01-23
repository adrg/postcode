package postcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestValidateByCountry(t *testing.T) {
	testCases := []struct {
		description string
		country     string
		postCode    string
		expected    bool
	}{
		{
			description: "Paris, France",
			country:     "FR",
			postCode:    "75008",
			expected:    true,
		},
		{
			description: "Brussels, Belgium",
			country:     "BE",
			postCode:    "1000",
			expected:    true,
		},
		{
			description: "Utrecht, The Netherlands",
			country:     "NL",
			postCode:    "3511 ax",
			expected:    true,
		},
		{
			description: "Utrecht, The Netherlands, Alt",
			country:     "NL",
			postCode:    "3511AX",
			expected:    true,
		},
		{
			description: "Hannover, Germany",
			country:     "DE",
			postCode:    "30179",
			expected:    true,
		},
		{
			description: "Vilnius, Lithuania",
			country:     "LT",
			postCode:    "LT-00200",
			expected:    true,
		},
		{
			description: "Empty postal code",
			country:     "FR",
			postCode:    "",
			expected:    false,
		},
		{
			description: "Short postal code",
			country:     "FR",
			postCode:    "A",
			expected:    false,
		},
		{
			description: "Non-existant country code",
			country:     "TY",
			postCode:    "TY 1234",
			expected:    false,
		},
		{
			description: "Non-existant postal code format",
			country:     "US",
			postCode:    "11111111111",
			expected:    false,
		},
		{
			description: "Postal code for wrong country",
			country:     "FR",
			postCode:    "1111",
			expected:    false,
		},
	}

	postCodeValidator := NewPostCodeValidator()

	for _, testCase := range testCases {
		isValid := postCodeValidator.ValidatePostalCode(testCase.country, testCase.postCode)
		assert.Equal(t, isValid, testCase.expected)
	}
}
