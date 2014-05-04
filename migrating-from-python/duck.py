#!/usr/bin/env python

class Cow(object):
    pass

class Fox(object):
    def get_noise(self):
        return "Ring-ding-ding-ding-dingeringeding"

def ask(obj):
    print('What goes "{}?"'.format(obj.get_noise()))

ask(Fox())
ask(Cow())
