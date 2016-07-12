postcode
========
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/adrg/postcode)
[![License: MIT](http://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](http://opensource.org/licenses/MIT)

Small package for validating postal codes. While the validation process does
not guarantee that the postcode actually exists, it does guarantee that the
format of the provided input is valid.

The reason for creating this package is that there is no good regular
expression for validating postal codes, and even if it existed, it would have
been huge and inefficient.

Full documentation can be found at: http://godoc.org/github.com/adrg/postcode

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

## License
Copyright (c) 2016 Adrian-George Bostan.

This project is licensed under the [MIT license](http://opensource.org/licenses/MIT). See LICENSE for more details.
