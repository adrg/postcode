package postcode

import (
	"testing"
)

func Test_Validate(t *testing.T) {
	testCases := []struct {
		description string
		input       string
		expected    error
	}{
		{
			description: "Valid postal code Paris France",
			input:       " 75008 ",
			expected:    nil,
		},
		{
			description: "Valid postal code Brussels Belgium",
			input:       " 1000 ",
			expected:    nil,
		},
		{
			description: "Valid postal code Utrecht The netherlands",
			input:       " 3511 AX ",
			expected:    nil,
		},
		{
			description: "Valid postal code Utrecht different format The Netherlands",
			input:       " 3511AX ",
			expected:    nil,
		},
		{
			description: "Valid postal code Hannover Germany",
			input:       " 30179 ",
			expected:    nil,
		},
		{
			description: "Invalid postal code empty",
			input:       "",
			expected:    ErrEmpty,
		},
		{
			description: "Invalid postal code short",
			input:       "A",
			expected:    ErrShort,
		},
		{
			description: "Invalid postal code invalid country",
			input:       "TY 1234",
			expected:    ErrInvalidCountry,
		},
		{
			description: "Invalid postal code invalid 11 chars length",
			input:       "11111111111",
			expected:    ErrInvalid,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.description, func(t *testing.T) {
			err := Validate(testCase.input)

			if err != testCase.expected {
				t.Errorf("expected: %v; got: %v", err, testCase.expected)
			}
		})
	}
}
