package challenge

/*
Package challenge find the 5 locations closests to the office
and the 5 locations furthests to the same point from a CSV file.

The CSV file must have a line structure like this

 	id, lat, lng
 	234123, 43.42321, -34.43322
where the first row is an Id and the others two are latitude and longitude
respectively. The first line with header can be omitted.

Usage: ./challenge <filename>
*/

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"

	"github.com/kellydunn/golang-geo"
)

// Point of reference to calculate distances.
var Office = geo.NewPoint(51.925146, 4.478617)

type Location struct {
	id       int64
	distance float64
}

// ParseId parse the id row from the CSV file.
func parseId(csvId string) int64 {
	var id int64
	var err error
	if id, err = strconv.ParseInt(csvId, 10, 0); err != nil {
		e := err.(*strconv.NumError)
		if e.Err != strconv.ErrSyntax {
			log.Fatal(e.Err)
		} else {
			return -1
		}
	}
	return id
}

// makePointFromCsv parse latitude and longitude from the rows in the CSV file
// and make a point from the library geo.
func makePointFromCsv(csvLat, csvLng string) *geo.Point {
	var lat, lng float64
	var err error
	if lat, err = strconv.ParseFloat(csvLat, 64); err != nil {
		e := err.(*strconv.NumError)
		if e.Err != strconv.ErrSyntax {
			log.Fatal(e.Err)
		}
	}
	if lng, err = strconv.ParseFloat(csvLng, 64); err != nil {
		e := err.(*strconv.NumError)
		if e.Err != strconv.ErrSyntax {
			log.Fatal(e.Err)
		}
	}
	point := geo.NewPoint(lat, lng)
	return point
}

// distanceToOffice calculate the distance, in kilometers from the *point* to the Office.
func distanceToOffice(point *geo.Point) float64 {
	return point.GreatCircleDistance(Office)
}

// makeClosestList make a ordered list with the closests points to the Office.
func makeClosestList(list []Location, location Location) []Location {
	index := sort.Search(
		len(list),
		func(i int) bool { return list[i].distance > location.distance },
	)
	if index < 5 {
		if len(list) < 5 {
			list = append(list, Location{})
		}
		copy(list[index+1:], list[index:])
		list[index] = location
	}
	return list
}

// makeFurthestList make a ordered list with the furthests points to the Office.
func makeFurthestList(list []Location, location Location) []Location {
	index := sort.Search(
		len(list),
		func(i int) bool { return list[i].distance < location.distance },
	)
	if index < 5 {
		if len(list) < 5 {
			list = append(list, Location{})
		}
		copy(list[index+1:], list[index:])
		list[index] = location
	}
	return list
}

// parseLocations open and read the CSV file and process every line to calculate
// the 5 closests and furthests points to the Office.
func parseLocations(filename string) ([]Location, []Location, error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	closests := make([]Location, 0, 5)
	furthests := make([]Location, 0, 5)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			return closests, furthests, nil
		} else if err != nil {
			log.Fatal(err)
		}
		// Parse the id
		id := parseId(line[0])
		if id == -1 {
			continue
		}
		// Make a point
		point := makePointFromCsv(line[1], line[2])
		// Calculate the distance to the Office
		distance := distanceToOffice(point)
		// Make the list of the closests and the furthests
		location := Location{
			id,
			distance,
		}

		closests = makeClosestList(closests, location)
		furthests = makeFurthestList(furthests, location)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Error: Source file name is required\n")
		fmt.Println("Usage:", os.Args[0], "<filename> \n")
		return
	}
	closests, furthests, err := parseLocations(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Closests")
	for _, points := range closests {
		fmt.Printf("ID: %d - Distance: %.2f Kms.\n", points.id, points.distance)
	}
	fmt.Println("Furthests")
	for _, points := range furthests {
		fmt.Printf("ID: %d - Distance: %.2f Kms.\n", points.id, points.distance)
	}
}
