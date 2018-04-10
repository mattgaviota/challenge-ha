# Challenge

We have some listings records in a csv file - one listing per line.
We would like to know what top 5 listings are closest (in terms of distance) to
our office (GPS coordinates - lat: 51.925146 lng: 4.478617) and what 5 listings
are the furthest.

You can use the first formula from this Wikipedia article (https://en.wikipedia.org/wiki/Great-circle_distance)
to calculate distance, don't forget, you'll need to convert degrees to radians.
Your program should be fully tested too.

Write a program that will read the full list of listings and output two top 5
lists with listings IDs and relative distances to our office

Your program should be fully tested too. Application should be written in Go and
core functionality should be covered by unit tests.
