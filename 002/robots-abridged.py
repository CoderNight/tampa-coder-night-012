#!/usr/bin/env python

# Algorithm:
# - Break up standard input into 4-line conveyor chunks tossing the 4th line
#
# - Find the location of X on the conveyor belt
# - Slice the northLasers starting from X in both directions keeping only
#   the even spaces
# - Do the same for southLasers, but starting one space to the left and
#   right of X so the "evens" will really be the odds.
# - Count the '|'s in the resulting slices
# - Go east if and only if east has fewer '|'s than west

from sys import stdin
input = stdin.readlines()
for i in range(0,len(input),4):
    northLasers, conveyorBelt, southLasers = input[i:i+3]
    x = conveyorBelt.index('X')
    west = (northLasers[x::-2] + southLasers[x-1::-2]).count('|')
    east = (northLasers[x::2]  + southLasers[x+1::2]).count('|')
    print 'GO EAST' if east < west else 'GO WEST'
