package postcode

import (
	"testing"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    error
	}{
		{
			description: "Paris_France_Valid",
			input:       " 75008 ",
			expected:    nil,
		},
		{
			description: "Brussels_Belgium_Valid",
			input:       " 1000 ",
			expected:    nil,
		},
		{
			description: "Utrecht_The_Netherlands_Valid",
			input:       " 3511 ax ",
			expected:    nil,
		},
		{
			description: "Utrecht_The_Netherlands_Alt_Valid",
			input:       " 3511AX ",
			expected:    nil,
		},
		{
			description: "Hannover_Germany_Valid",
			input:       " 30179 ",
			expected:    nil,
		},
		{
			description: "Empty_Postal_Code_Invalid",
			input:       "",
			expected:    ErrEmpty,
		},
		{
			description: "Short_Postal_Code_Invalid",
			input:       "A",
			expected:    ErrShort,
		},
		{
			description: "Inexistent_Country_Code_Invalid",
			input:       "TY 1234",
			expected:    ErrInvalidCountry,
		},
		{
			description: "Inexistent_Postal_Code_Format_Invalid",
			input:       "11111111111",
			expected:    ErrInvalidFormat,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			if err := Validate(testCase.input); err != testCase.expected {
				t.Errorf("expected: %v; got: %v", err, testCase.expected)
			}
		})
	}
}
