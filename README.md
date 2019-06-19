# eurex

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
