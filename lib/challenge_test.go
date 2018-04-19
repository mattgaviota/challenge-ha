package lib

import (
	"testing"

	"github.com/kellydunn/golang-geo"
)

func TestParseId(t *testing.T) {
	cases := []struct {
		in   string
		want int64
	}{
		{"23423", 23423},
		{"id", -1},
		{"", -1},
		{"8917237612736871236712637123", -2},
	}
	for _, c := range cases {
		got := ParseId(c.in)
		if got != c.want {
			t.Errorf("ParseId(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
}

func TestMakePointFromCsv(t *testing.T) {
	point1 := geo.NewPoint(37.1768672, -3.608897)
	point2 := geo.NewPoint(-1, -1)
	point3 := geo.NewPoint(-2, -2)
	cases := []struct {
		in1  string
		in2  string
		want *geo.Point
	}{
		{"37.1768672", "-3.608897", point1},
		{"lat", "lng", point2},
		{"27.454334", "lng", point2},
		{"16384e1024", "34.567", point3},
		{"27.431414", "16384e1024", point3},
	}
	for _, c := range cases {
		got := MakePointFromCsv(c.in1, c.in2)
		if *got != *c.want {
			t.Errorf("MakePointFromCsv(%q, %q) == %#v, want %#v", c.in1, c.in2, got, c.want)
		}
	}
}

func TestDistanceToOffice(t *testing.T) {
	point := geo.NewPoint(37.1768672, -3.608897)
	cases := []struct {
		in   *geo.Point
		want float64
	}{
		{point, 1758.0806131200134},
	}
	for _, c := range cases {
		got := DistanceToOffice(c.in)
		if got != c.want {
			t.Errorf("DistanceToOffice(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
}

func testEq(a, b []Location) bool {
	if &a == &b {
		return true
	}
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Id != b[i].Id {
			return false
		}
		if a[i].Distance != b[i].Distance {
			return false
		}
	}
	return true
}

func TestMakeClosestList(t *testing.T) {
	location1 := Location{3212, 0.53123}
	location2 := Location{8976, 0.74567}
	location3 := Location{2456, 0.85421}
	emptyList := []Location{}
	oneElementList := []Location{
		location1,
	}
	twoElementList := []Location{
		location1,
		location3,
	}
	threeElementList := []Location{
		location1,
		location2,
		location3,
	}
	amount := 3
	cases := []struct {
		inList     []Location
		inLocation Location
		want       []Location
	}{
		{emptyList, location1, oneElementList},
		{twoElementList, location2, threeElementList},
	}
	for _, c := range cases {
		got := MakeClosestList(c.inList, c.inLocation, amount)
		if !testEq(got, c.want) {
			t.Errorf("MakeClosestList(%#v, %#v, %#v) == %#v, want %#v", c.inList, c.inLocation, amount, got, c.want)
		}
	}
}

func TestMakeFurthestList(t *testing.T) {
	location1 := Location{2456, 0.85421}
	location2 := Location{8976, 0.74567}
	location3 := Location{3212, 0.53123}
	emptyList := []Location{}
	oneElementList := []Location{
		location1,
	}
	twoElementList := []Location{
		location1,
		location3,
	}
	threeElementList := []Location{
		location1,
		location2,
		location3,
	}
	amount := 3
	cases := []struct {
		inList     []Location
		inLocation Location
		want       []Location
	}{
		{emptyList, location1, oneElementList},
		{twoElementList, location2, threeElementList},
	}
	for _, c := range cases {
		got := MakeFurthestList(c.inList, c.inLocation, amount)
		if !testEq(got, c.want) {
			t.Errorf("MakeFurthestList(%#v, %#v, %#v) == %#v, want %#v", c.inList, c.inLocation, amount, got, c.want)
		}
	}
}
