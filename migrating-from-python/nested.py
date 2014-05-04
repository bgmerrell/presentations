#!/usr/bin/env python

class C(object):
    def __init__(self, filename):
        self.filename = filename
    def printfile(self):
        try:
            with open(self.filename) as f:
                print f.read()
        except IOError:
            print("Failed to open file")

c = C("/var/tmp/file.txt")
c.printfile()
