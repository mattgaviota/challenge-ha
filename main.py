#!/usr/bin/env python
# -*- coding: utf-8 -*-

import sys
from calculator import DistanceCalculator


def main(filename):
    calculator = DistanceCalculator(filename)
    closests = calculator.closests(5)
    furthests = calculator.furthests(5)
    print('Closests: ')
    print(closests)
    print('furthests: ')
    print(furthests)

if __name__ == '__main__':
    filename = sys.argv[1]
    main(filename)
