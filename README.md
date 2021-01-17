postcode
========
[![Build Status](https://github.com/adrg/postcode/workflows/CI/badge.svg)](https://github.com/adrg/postcode/actions?query=workflow%3ACI)
[![Code coverage](https://codecov.io/gh/adrg/postcode/branch/master/graphs/badge.svg?branch=master)](https://codecov.io/gh/adrg/postcode)
[![pkg.go.dev documentation](https://pkg.go.dev/badge/github.com/adrg/postcode)](https://pkg.go.dev/github.com/adrg/postcode)
[![MIT License](https://img.shields.io/badge/license-MIT-red.svg?style=flat-square)](https://opensource.org/licenses/MIT)
[![Go report card](https://goreportcard.com/badge/github.com/adrg/postcode)](https://goreportcard.com/report/github.com/adrg/postcode)
[![GitHub issues](https://img.shields.io/github/issues/adrg/postcode)](https://github.com/adrg/postcode/issues)
[![Buy me a coffee](https://img.shields.io/static/v1.svg?label=%20&message=Buy%20me%20a%20coffee&color=579fbf&logo=buy%20me%20a%20coffee&logoColor=white)](https://ko-fi.com/T6T72WATK)

Small package for validating postal codes. While the validation process does
not guarantee that the postcode actually exists, it does guarantee that the
format of the provided input is valid.

The reason for creating this package is that there is no good regular
expression for validating postal codes, and even if it existed, it would have
been huge and inefficient.

Full documentation can be found at: https://pkg.go.dev/github.com/adrg/postcode.

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
        // Treat error.
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
See [CONTRIBUTING.MD](CONTRIBUTING.md).

## License
Copyright (c) 2016 Adrian-George Bostan.

This project is licensed under the [MIT license](https://opensource.org/licenses/MIT).
See [LICENSE](LICENSE) for more details.
