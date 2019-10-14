postcode
========
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/adrg/postcode)
[![License: MIT](https://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/adrg/postcode)](https://goreportcard.com/report/github.com/adrg/postcode)

Small package for validating postal codes. While the validation process does
not guarantee that the postcode actually exists, it does guarantee that the
format of the provided input is valid.

The reason for creating this package is that there is no good regular
expression for validating postal codes, and even if it existed, it would have
been huge and inefficient.

Full documentation can be found at: https://godoc.org/github.com/adrg/postcode

## Installation
    go get github.com/adrg/postcode

## Usage

```go
package main

import (
	"github.com/adrg/postcode"
)

func main() {
	if err := postcode.Validate("10007"); err != nil {
        // the postcode is not valid
        // treat error
	}
}
```

## References
For more information see
* [Wikipedia's List of postal codes](https://en.wikipedia.org/wiki/List_of_postal_codes)
* [Wikipedia's List of country codes](https://en.wikipedia.org/wiki/ISO_3166-1)

## Contributing

Contributions in the form of pull requests, issues or just general feedback,
are always welcome.
See [CONTRIBUTING.MD](https://github.com/adrg/postcode/blob/master/CONTRIBUTING.md).

## License
Copyright (c) 2016 Adrian-George Bostan.

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).
See [LICENSE](https://github.com/adrg/postcode/blob/master/LICENSE) for more details.
