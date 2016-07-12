package postcode

import (
	"errors"
	"strings"
	"unicode"
)

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
