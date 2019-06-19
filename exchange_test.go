package eurex

import (
	"fmt"
	"testing"
)

// TestGet will test the Get function
func TestGet(t *testing.T) {
	var tests = []struct {
		quantity      float64
		currencyFrom  string
		currencyTo    string
		time          string
		errorExpected string
	}{
		{
			quantity:      0,
			currencyFrom:  "",
			currencyTo:    "",
			time:          "",
			errorExpected: "quantity must be more than 0",
		},
		{
			quantity:      2,
			currencyFrom:  "",
			currencyTo:    "",
			time:          "",
			errorExpected: "currencyFrom code must be passed",
		},
		{
			quantity:      2,
			currencyFrom:  "USD",
			currencyTo:    "",
			time:          "",
			errorExpected: "currencyTo code must be passed",
		},
		{
			quantity:      2,
			currencyFrom:  "USD",
			currencyTo:    "CHF",
			time:          "",
			errorExpected: "time code must be passed",
		},
		{
			quantity:      2,
			currencyFrom:  "Code not valid",
			currencyTo:    "CHF",
			time:          "2019-06-12",
			errorExpected: "currency code (\"Code not valid\") not found",
		},
		{
			quantity:      2,
			currencyFrom:  "USD",
			currencyTo:    "Code not valid",
			time:          "2019-06-12",
			errorExpected: "currency code (\"Code not valid\") not found",
		},
	}

	// Init the client
	client := NewClient()

	for i, test := range tests {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			_, err := client.Exchange.Get(test.quantity, test.currencyFrom, test.currencyTo, test.time)
			if err != nil && err.Error() != test.errorExpected {
				t.Errorf("Test failed: wrong error\nwant: %q\ngot:  %q", test.errorExpected, err.Error())
			}
			if err == nil && test.errorExpected != "" {
				t.Errorf("Test failed: no error\nexpected:  %q", test.errorExpected)
			}
		})
	}
}
