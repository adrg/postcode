package postcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestValidateByCountry(t *testing.T) {
	testCases := []struct {
		description   string
		country       string
		postCode      string
		expected      bool
		errorExpected bool
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
			description:   "Non-existant country code",
			country:       "TY",
			postCode:      "TY 1234",
			expected:      false,
			errorExpected: true,
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
		{
			description: "UK Postcode",
			country:     "GB",
			postCode:    "KT111AT",
			expected:    true,
		},
	}

	postCodeValidator := NewPostCodeValidator()

	for _, testCase := range testCases {
		t.Run(testCase.description, func(tt *testing.T) {
			isValid, err := postCodeValidator.ValidatePostalCode(testCase.country, testCase.postCode)
			if !testCase.errorExpected {
				assert.NilError(tt, err)
			}
			assert.Equal(tt, isValid, testCase.expected)
		})
	}
}
