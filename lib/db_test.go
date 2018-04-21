package lib

import (
	"testing"
)

func TestQueryLocations(t *testing.T) {
	driver := "postgres"
	wrongDriver := "mysql"
	datasource := "postgres://admin:admin@postgres/challengedb?sslmode=disable"
	table := "locations"
	wrongTable := "points"
	amount := 3
	closestList := []Location{
		{442406, 0.3338381556847423},
		{285782, 0.5280320575225937},
		{429151, 0.6480104060043963},
	}
	furthestList := []Location{
		{7818, 8776.646278220525},
		{382013, 1810.117975607402},
		{381823, 1758.8482704875876},
	}
	cases := []struct {
		inDriver         string
		inDataSource     string
		inTable          string
		inAmount         int
		wantClosestList  []Location
		wantFurthestList []Location
		wantError        int
	}{
		{driver, datasource, table, amount, closestList, furthestList, 0}, // All ok
		{wrongDriver, datasource, table, amount, nil, nil, 201},           // Wrong dataSource
		{driver, datasource, wrongTable, amount, nil, nil, 202},           // Wrong table
	}
	for _, c := range cases {
		gotClosestList, gotFurthestList, gotErr := QueryLocations(c.inDriver, c.inDataSource, c.inTable, c.inAmount)
		if !testEq(gotClosestList, c.wantClosestList) || !testEq(gotFurthestList, c.wantFurthestList) || gotErr != c.wantError {
			t.Errorf(
				"Query(%#v, %#v, %#v, %#v) == %#v %#v %#v want %#v %#v %#v",
				c.inDriver,
				c.inDataSource,
				c.inTable,
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
