# Challenge

## Description

Simple package to find the 5 locations closests to the office
and the 5 locations furthests to the same point from a CSV file.

The CSV file must have a line structure like this

 	id, lat, lng
 	234123, 43.42321, -34.43322

where the first row is an Id and the others two are latitude and longitude
respectively. The first line with header can be omitted.

The amount of locations can be changed with the amount argument. Default is 5.

Also you can use a table from a postgres database with the same structure of the
CSV file.

## Requirements

In order to try the app you need [golang](https://golang.org/doc/install),
the [geo library](https://github.com/kellydunn/golang-geo), the [postgres driver](github.com/lib/pq) and [argparse](github.com/akamensky/argparse)

You can install all the libraries with:

    go get https://github.com/kellydunn/golang-geo
    go get github.com/lib/pq
    go get github.com/akamensky/argparse

## Try it

You need install the package first. First download or clone the folder and next

    cd challenge-ha/challenge && go install .
    cd $GOPATH/bin
    ./challenge -f <PATHTOPACKAGE>/challenge-ha/testdata/geoDataLarge.csv -a 5

If you want to test using a database, you can use [docker](https://www.docker.com/community-edition#/download) and [docker-compose](https://docs.docker.com/compose/install/)

    docker-compose up -d
    ./challenge -d postgres -s "postgres://admin:admin@postgres/challengedb?sslmode=disable" -t locations -a 5

## Usage

    usage: Challenge [-h|--help] [-f|--filename "<value>"] [-d|--driver "<value>"]
                     [-s|--datasource "<value>"] [-t|--table "<value>"]
                     [-a|--amount <integer>]

                     Challenge cli

    Arguments:

      -h  --help        Print help information
      -f  --filename    filename of the CSV(Ex. path/to/file.csv)
      -d  --driver      driver of the database(Ex. 'postgres'). Default: postgres
      -s  --datasource  datasource of the database(Ex.
                        'postgres://user:pass@host/database')
      -t  --table       table where the id, lat and lng are stored(Ex.
                        'locations'). Default: locations
      -a  --amount      Amount of locations included in the lists. Default: 5

## Test

You can run the tests with:

    cd lib/
    go test

You should setup the docker container with the database to complete all the tests.
