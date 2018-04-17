# Challenge

## Description

Simple package to find the 5 locations closests to the office
and the 5 locations furthests to the same point from a CSV file.

The CSV file must have a line structure like this

 	id, lat, lng
 	234123, 43.42321, -34.43322

where the first row is an Id and the others two are latitude and longitude
respectively. The first line with header can be omitted.

## Requirements

In order to try the package you need golang(https://golang.org/doc/install) and
the geo library(https://github.com/kellydunn/golang-geo)

You can install geo library with:

    go get https://github.com/kellydunn/golang-geo

## Usage

You need install the package first. First download or clone the folder and next

    cd challenge-ha/challenge && go install .
    cd $GOPATH/bin
    ./challenge <PATHTOPACKAGE>/challenge-ha/files/geoData.csv
