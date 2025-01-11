package main

import (
	"fmt"
	"os"

	"github.com/13excite/c24-expences/pkg/c24parser"
)

func main() {
	csvParser := c24parser.NewParser()
	if err := csvParser.ParseFile("input/transaction_12_24.csv"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Transactions:")
	for _, t := range csvParser.GetTransactions() {
		fmt.Printf("%+v\n", t)
	}
}
