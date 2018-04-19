package main

/*
Package main find the amount locations from a set of data. This data can be a
CSV file or a table from a DB.

Usage:

	./challenge -f <filename> [amount]
or
	./challenge -d -h localhost -u user -p password -t table amount
*/

import (
	"challenge-ha/lib"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Error: Source file name is required. Amount locations is optional\n")
		fmt.Println("Usage:", os.Args[0], "<filename> [amount]")
		return
	}
	amount := 5
	if len(os.Args) > 2 {
		var err error
		if amount, err = strconv.Atoi(os.Args[2]); err != nil {
			amount = 5
		}
	}
	closests, furthests, err := lib.ParseLocations(os.Args[1], amount)
	if err != 0 {
		switch err {
		case -1:
			fmt.Println("The file doesn't exist")
		case -2:
			fmt.Println("The CSV file has a line with the wrong number of fields")
		default:
			fmt.Println(err)
		}
	} else {
		fmt.Println("Closests")
		for _, points := range closests {
			fmt.Printf("ID: %d - Distance: %.2f Kms.\n", points.Id, points.Distance)
		}
		fmt.Println("Furthests")
		for _, points := range furthests {
			fmt.Printf("ID: %d - Distance: %.2f Kms.\n", points.Id, points.Distance)
		}
	}
}
