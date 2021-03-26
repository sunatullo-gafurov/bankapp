package main

import (
	"bank/pkg/bank/deposit"
	"fmt"
)

func main() {
	min, max := deposit.Calculate(5_000_00, "USD")
	fmt.Println(min, max)
}
