package main

/*
Package main find the amount locations from a set of data. This data can be a
CSV file or a table from a DB.

Usage:

usage: Challenge [-h|--help] [-f|--filename "<value>"] [-d|--datasource
                 "<value>"] [-a|--amount <integer>]

                 Challenge program

Arguments:

  -h  --help        Print help information
  -f  --filename    filename of the CSV(Ex. path/to/file.csv)
  -d  --datasource  datasource of the table(Ex.
                    'postgres://user:pass@host/database')
  -a  --amount      Amount of locations included in the lists. Default: 5
*/

import (
	"challenge-ha/lib"
	"fmt"
	"os"

	"github.com/akamensky/argparse"
)

func showResult(closests []lib.Location, furthests []lib.Location, err int) {
	if err != 0 {
		lib.ShowError(err)
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
	os.Exit(0)
}

func main() {
	parser := argparse.NewParser("Challenge", "Challenge cli")

	// Optional shorthand argument
	filename := parser.String("f", "filename", &argparse.Options{Required: false, Help: "filename of the CSV(Ex. path/to/file.csv)"})
	// Optional shorthand argument
	datasource := parser.String("d", "datasource", &argparse.Options{Required: false, Help: "datasource of the table(Ex. 'postgres://user:pass@host/database')"})
	// Optional shorthand argument
	amount := parser.Int("a", "amount", &argparse.Options{Required: false, Default: 5, Help: "Amount of locations included in the lists"})
	// Parse args
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	if *filename != "" {
		fmt.Println("Results calculated from CSV File")
		closests, furthests, err := lib.ParseLocations(*filename, *amount)
		showResult(closests, furthests, err)
	}
	if *datasource != "" {
		fmt.Println("Results calculated from database table")
		closests, furthests, err := lib.QueryLocations(*datasource, *amount)
		showResult(closests, furthests, err)
	}
	fmt.Print(parser.Usage(err))
	os.Exit(1)
}
