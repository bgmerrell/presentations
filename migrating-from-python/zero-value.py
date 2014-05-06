#!/usr/bin/env python

def is_len_valid(foo):
	return bool(len(foo) % 2)

print is_len_valid([1,2,3,4])
print is_len_valid(None)
