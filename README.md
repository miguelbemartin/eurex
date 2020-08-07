# eurex

This Go Library allows you to retrieve currency exchange rates from the official [European Central Bank service](https://www.ecb.europa.eu). 

[![Go Report Card](https://goreportcard.com/badge/github.com/miguelbemartin/eurex)](https://goreportcard.com/report/github.com/miguelbemartin/eurex)
[![Build Status](https://travis-ci.org/miguelbemartin/eurex.svg?branch=master)](https://travis-ci.org/miguelbemartin/eurex)
[![codecov](https://codecov.io/gh/miguelbemartin/eurex/branch/master/graph/badge.svg)](https://codecov.io/gh/miguelbemartin/eurex)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/miguelbemartin/eurex/master/LICENSE)

## Getting Started

### Installation

```
go get github.com/miguelbemartin/eurex
```

### Requirements

- Go 1.10+

### Usage

```go

//...

import (
	eurex "github.com/miguelbemartin/eurex"
)

//...

client := eurex.NewClient()

rate, err := client.Exchange.Get(14.50, "USD", "CHF", "2019-06-15")
if err != nil {
  // handle your error
}
```

### Run tests

```
go test . -v
```

## Contributing
* Open a PR: https://github.com/miguelbemartin/eurex/pulls
* Open an issue: https://github.com/miguelbemartin/eurex/issues

## Authors
* **Miguel Ángel Martín** - [@miguelbemartin](https://twitter.com/miguelbemartin)

## License
This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
