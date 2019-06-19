# eurex

A Go library to handle Currency exchange rates

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
