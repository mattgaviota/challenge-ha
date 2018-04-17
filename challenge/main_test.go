package main

import (
	"testing"

	geo "github.com/kellydunn/golang-geo"
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
		got := parseId(c.in)
		if got != c.want {
			t.Errorf("parseId(%#v) == %#v, want %#v", c.in, got, c.want)
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
		got := makePointFromCsv(c.in1, c.in2)
		if *got != *c.want {
			t.Errorf("makePointFromCsv(%q, %q) == %#v, want %#v", c.in1, c.in2, got, c.want)
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
		got := distanceToOffice(c.in)
		if got != c.want {
			t.Errorf("distanceToOffice(%#v) == %#v, want %#v", c.in, got, c.want)
		}
	}
}
