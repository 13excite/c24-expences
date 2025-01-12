package main

import (
	"fmt"
	"os"

	"github.com/13excite/c24-expences/pkg/c24parser"
	"github.com/13excite/c24-expences/pkg/driver"
	"github.com/13excite/c24-expences/pkg/models"
)

func main() {

	conn, err := driver.OpenDB("default", "", "localhost:9000", "default")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()

	model := models.NewModels(conn)

	csvParser := c24parser.NewParser()
	if err := csvParser.ParseFile("input/transaction_12_24.csv"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Transactions:")
	for _, t := range csvParser.GetTransactions() {
		err := model.DB.InsertTransaction(t)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
