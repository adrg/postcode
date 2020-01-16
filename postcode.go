/*
Package postcode validates postal codes. While the validation process does
not guarantee that the postcode actually exists, it does guarantee that the
format of the provided input is valid.

	For more information see:
	https://en.wikipedia.org/wiki/List_of_postal_codes
	https://en.wikipedia.org/wiki/ISO_3166-1

Example

	if err := postcode.Validate("10007"); err != nil {
		// the postcode is not valid
		// treat error
	}
*/
package postcode

import (
	"errors"
	"strings"
	"unicode"
)

// Errors.
var (
	ErrEmpty          = errors.New("postal code cannot be empty")
	ErrShort          = errors.New("postal code cannot be shorter than 2 characters")
	ErrInvalidCountry = errors.New("invalid country code")
	ErrInvalid        = errors.New("invalid postal code format")
)

// Validate checks if the provided input string matches any of the
// accepted postcode formats. If the validation fails, the function returns
// an error specifying the cause.
func Validate(code string) error {
	if code = strings.ToUpper(strings.TrimSpace(code)); code == "" {
		return ErrEmpty
	}

	format := []rune(code)
	if len(format) < 2 {
		return ErrShort
	}

	countryCode := string(format[:2])
	for i, r := range format {
		switch {
		case unicode.IsDigit(r):
			r = 'N'
		case unicode.IsLetter(r):
			r = 'A'
		}

		format[i] = r
	}

	if !inSlice(string(format), validFormats) {
		if format[0] == 'A' && format[1] == 'A' {
			format[0], format[1] = 'C', 'C'
			if inSlice(string(format), validFormats) {
				if !inSlice(countryCode, countryCodes) {
					return ErrInvalidCountry
				}

				return nil
			}
		}

		return ErrInvalid
	}

	return nil
}
