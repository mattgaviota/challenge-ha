# -*- coding: utf-8 -*-

from csv import DictReader
from geopy.distance import great_circle
from operator import itemgetter
from utils import TopList


OFFICE = (51.925146, 4.478617)

class DistanceCalculator(object):
    """ Distance calculator"""
    def __init__(self, filename):
        self.distances = []
        csvfile = open(filename)
        self.reader = DictReader(csvfile)
        self.calculateDistances()

    def closests(self, amount=5):
        """Find the 'amount' closest places to OFFICE."""
        return TopList(self.distances[:amount])

    def furthests(self, amount=5):
        """Resturn the 'amount' furthest places to OFFICE."""
        return TopList(self.distances[-amount:])

    def calculateDistances(self):
        """Find the distances between each item in the file and the OFFICE."""
        for row in self.reader:
            id = row['id']
            distance = great_circle(OFFICE, (row['lat'], row['lng'])).km
            self.distances.append((id, distance))
        self.distances.sort(key=itemgetter(1))
