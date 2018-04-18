package lib

/*
Library csv find the 5 locations closests to the office
and the 5 locations furthests to the same point from a CSV file.

The CSV file must have a line structure like this

 	id, lat, lng
 	234123, 43.42321, -34.43322
where the first row is an Id and the others two are latitude and longitude
respectively. The first line with header can be omitted.

The amount of locations can be changed with the amount argument. Default is 5.
*/
import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

// ParseLocations open and read the CSV file and process every line to calculate
// the 5 closests and furthests points to the Office.
func ParseLocations(filename string, amount int) ([]Location, []Location, error) {
	csvFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()
	reader := csv.NewReader(csvFile)
	closests := make([]Location, 0, amount)
	furthests := make([]Location, 0, amount)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			return closests, furthests, nil
		} else if err != nil {
			log.Fatal(err)
		}
		// Parse the id
		id := ParseId(line[0])
		if id == -1 {
			continue
		}
		// Make a point
		point := MakePointFromCsv(line[1], line[2])
		// Calculate the distance to the Office
		distance := DistanceToOffice(point)
		// Make the list of the closests and the furthests
		location := Location{
			id,
			distance,
		}

		closests = MakeClosestList(closests, location, amount)
		furthests = MakeFurthestList(furthests, location, amount)
	}
}
