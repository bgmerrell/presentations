#!/usr/bin/env python

import random

class C(object):
	def __setattr__(self, name, value):
		pass
	def __getattr__(self, name):
		return random.randint(100, 200)

# START OMIT
c = C()
c.foo = 42
print c.foo
# END OMIT
