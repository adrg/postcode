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
			description: "Inexistent country code",
			input:       "TY 1234",
			expected:    ErrInvalidCountry,
		},
		{
			description: "Inexistent postal code format",
			input:       "11111111111",
			expected:    ErrInvalidFormat,
		},
	}

	for _, testCase := range testCases {
		t.Logf("Validating `%s` (%s)", testCase.input, testCase.description)
		if err := Validate(testCase.input); err != testCase.expected {
			t.Errorf("expected: %v; got: %v", err, testCase.expected)
		}
	}
}
