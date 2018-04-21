package lib

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func QueryLocations(dataSource string, amount int) ([]Location, []Location, int) {
	db, err := sql.Open("postgres", dataSource)
	if err != nil {
		log.Print(err)
		return nil, nil, 201 // Error
	}

	rows, err := db.Query("SELECT * FROM locations")
	if err != nil {
		log.Print(err)
		return nil, nil, 202 // Error
	}
	defer rows.Close()
	closests := make([]Location, 0, amount)
	furthests := make([]Location, 0, amount)

	for rows.Next() {
		var id int64
		var lat, lng string
		err := rows.Scan(&id, &lat, &lng)
		if err != nil {
			log.Print(err)
			continue
		}

		// Make a point
		point := MakePointFromCsv(lat, lng)
		// Calculate the distance to the Office
		distance := DistanceToOffice(point)
		// Make the list of the closests and the furthests
		location := Location{id, distance}

		closests = MakeClosestList(closests, location, amount)
		furthests = MakeFurthestList(furthests, location, amount)
	}

	if err = rows.Err(); err != nil {
		log.Print(err)
		return nil, nil, 203 // Error reading rows
	}

	return closests, furthests, 0
}
