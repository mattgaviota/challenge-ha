package lib

import (
	"log"
	"sort"
	"strconv"

	"github.com/kellydunn/golang-geo"
)

// Point of reference to calculate distances.
var Office = geo.NewPoint(51.925146, 4.478617)

type Location struct {
	Id       int64
	Distance float64
}

// ParseId parse the id row from the CSV file.
func ParseId(csvId string) int64 {
	var id int64
	var err error
	if id, err = strconv.ParseInt(csvId, 10, 0); err != nil {
		e := err.(*strconv.NumError)
		if e.Err != strconv.ErrSyntax {
			log.Print(e.Err)
			return -2
		} else {
			return -1
		}
	}
	return id
}

// MakePointFromCsv parse latitude and longitude from the rows in the CSV file
// and make a point from the library geo.
func MakePointFromCsv(csvLat, csvLng string) *geo.Point {
	var lat, lng float64
	var err error
	if lat, err = strconv.ParseFloat(csvLat, 64); err != nil {
		e := err.(*strconv.NumError)
		if e.Err != strconv.ErrSyntax {
			log.Print(e.Err)
			return geo.NewPoint(-2, -2)
		} else {
			return geo.NewPoint(-1, -1)
		}
	}
	if lng, err = strconv.ParseFloat(csvLng, 64); err != nil {
		e := err.(*strconv.NumError)
		if e.Err != strconv.ErrSyntax {
			log.Print(e.Err)
			return geo.NewPoint(-2, -2)
		} else {
			return geo.NewPoint(-1, -1)
		}
	}
	return geo.NewPoint(lat, lng)
}

// DistanceToOffice calculate the distance, in kilometers from the *point* to the Office.
func DistanceToOffice(point *geo.Point) float64 {
	return point.GreatCircleDistance(Office)
}

// MakeClosestList make a ordered list with the closests points to the Office.
func MakeClosestList(list []Location, location Location, amount int) []Location {
	index := sort.Search(
		len(list),
		func(i int) bool { return list[i].Distance > location.Distance },
	)
	if index < amount {
		if len(list) < amount {
			list = append(list, Location{})
		}
		copy(list[index+1:], list[index:])
		list[index] = location
	}
	return list
}

// MakeFurthestList make a ordered list with the furthests points to the Office.
func MakeFurthestList(list []Location, location Location, amount int) []Location {
	index := sort.Search(
		len(list),
		func(i int) bool { return list[i].Distance < location.Distance },
	)
	if index < amount {
		if len(list) < amount {
			list = append(list, Location{})
		}
		copy(list[index+1:], list[index:])
		list[index] = location
	}
	return list
}
