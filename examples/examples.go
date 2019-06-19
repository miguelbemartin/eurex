package main

import (
	"fmt"

	"github.com/miguelbemartin/eurex"
)

func main() {
	fmt.Println("Running the example")

	client := eurex.NewClient()

	rsp, err := client.Exchange.Get(14.50, "USD", "CHF", "2019-06-13")
	if err != nil {
		fmt.Println("Error found: ", err.Error())
	} else {
		fmt.Println("Result:", *rsp)
	}

}
