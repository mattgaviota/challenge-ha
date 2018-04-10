# -*- coding: utf-8 -*-


class TopList(object):
    """Regular list with a personalized STR."""
    def __init__(self, items):
        self.items = items

    def __str__(self):
        return '\n'.join('id: {}, Distance: {:.2f} Kms.'.format(*t) for t in self.items)
