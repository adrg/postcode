package postcode

import (
	"testing"

	"gotest.tools/assert"
)

func TestValidate(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    error
	}{
		{
			description: "Paris, France",
			input:       " 75008 ",
			expected:    nil,
		},
		{
			description: "Brussels, Belgium",
			input:       " 1000 ",
			expected:    nil,
		},
		{
			description: "Utrecht, The Netherlands",
			input:       " 3511 ax ",
			expected:    nil,
		},
		{
			description: "Utrecht, The Netherlands, Alt",
			input:       " 3511AX ",
			expected:    nil,
		},
		{
			description: "Hannover, Germany",
			input:       " 30179 ",
			expected:    nil,
		},
		{
			description: "Vilnius, Lithuania",
			input:       "LT-00200",
			expected:    nil,
		},
		{
			description: "Empty postal code",
			input:       "",
			expected:    ErrEmpty,
		},
		{
			description: "Short postal code",
			input:       "A",
			expected:    ErrShort,
		},
		{
			description: "Non-existant country code",
			input:       "TY 1234",
			expected:    ErrInvalidCountry,
		},
		{
			description: "Non-existant postal code format",
			input:       "11111111111",
			expected:    ErrInvalidFormat,
		},
	}

	for _, testCase := range testCases {
		assert.Equal(t, Validate(testCase.input), testCase.expected)
	}
}
