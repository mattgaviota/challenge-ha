package lib

import (
	"testing"
)

func Test(t *testing.T) {
	filename := "../testdata/geoData.csv"
	wrongFilename := "../testdata/malformedData.csv"
	nonExistFilename := "../testdata/wrongName.csv"
	amount := 3
	closestList := []Location{
		{28403, 0.9222312128574133},
		{436807, 2.4448331880502057},
		{41238, 2.7057356681445883},
	}
	furthestList := []Location{
		{7818, 8776.646278220525},
		{382582, 1758.0806131200134},
		{419117, 371.313203576673},
	}
	cases := []struct {
		inFilename       string
		inAmount         int
		wantClosestList  []Location
		wantFurthestList []Location
		wantError        int
	}{
		{filename, amount, closestList, furthestList, 0}, // All ok
		{nonExistFilename, amount, nil, nil, 101},        // File doesn't exist
		{wrongFilename, amount, nil, nil, 102},           // Malformed CSV
	}
	for _, c := range cases {
		gotClosestList, gotFurthestList, gotErr := ParseLocations(c.inFilename, c.inAmount)
		if !testEq(gotClosestList, c.wantClosestList) || !testEq(gotFurthestList, c.wantFurthestList) || gotErr != c.wantError {
			t.Errorf(
				"ParseLocations(%#v, %#v) == %#v %#v %#v want %#v %#v %#v",
				c.inFilename,
				c.inAmount,
				gotClosestList,
				gotFurthestList,
				gotErr,
				c.wantClosestList,
				c.wantFurthestList,
				c.wantError,
			)
		}
	}
}
