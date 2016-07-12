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

// Validate checks if the provided input string matches any of the
// accepted postcode formats. If the validation fails, the function returns
// an error specifying the cause.
func Validate(code string) error {
	if code = strings.ToUpper(strings.TrimSpace(code)); code == "" {
		return errors.New("Postal code cannot be empty")
	}

	rCode := []rune(code)
	if len(rCode) < 2 {
		return errors.New("Postal code cannot be shorter than 2 characters")
	}

	countryCode := string(rCode[:2])
	for i, rc := range rCode {
		r := rc
		switch {
		case unicode.IsDigit(r):
			r = 'N'
		case unicode.IsLetter(r):
			r = 'A'
		}

		rCode[i] = r
	}

	if !inSlice(string(rCode), validFormats) {
		if rCode[0] == 'A' && rCode[1] == 'A' {
			rCode[0], rCode[1] = 'C', 'C'
			if inSlice(string(rCode), validFormats) {
				if !inSlice(countryCode, countryCodes) {
					return errors.New("Invalid country code")
				}

				return nil
			}
		}

		return errors.New("Invalid postal code format")
	}

	return nil
}
