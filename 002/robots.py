#!/usr/bin/env python

"""
Solution for the Robots vs Lasers problem:
http://www.puzzlenode.com/puzzles/4-robots-vs-lasers

If run directly, accepts a list of conveyors on standard input.

=== Sample input ===
#|#|#|##
---X----
###||###

##|#|#|#
----X---
###||###

##|#|#|#
----X---
###|||##
====================

=== Sample output ==
GO WEST
GO EAST
GO WEST
====================
"""

def bestDirectionForLeastDamage(northLasers, conveyorBelt, southLasers):
    """
    Given the north and south side laser maps of a conveyor,
    return the best direction for X in conveyor to travel
    for the least amount of damage where the north side fires
    only on X's even-numbered steps and the south side fires
    on X's odd-numbered steps.
    """
    # Algorithm:
    # - Find the location of X on the conveyor belt
    # - Slice the northLasers starting from X in both directions keeping only
    #   the even spaces
    # - Do the same for southLasers, but starting one space to the left and of X
    #   so the "evens" will really be the odds.
    # - Count the '|'s in the resulting slices
    # - Go east if and only if east has fewer '|'s than west
    x = conveyorBelt.index('X')
    west = (northLasers[x::-2] + southLasers[x-1::-2]).count('|')
    east = (northLasers[x::2]  + southLasers[x+1::2]).count('|')
    return 'GO EAST' if east < west else 'GO WEST'

def bestDirectionsForLeastDamage(conveyors):
    """
    Return a list of directional results, east or west, for a given list
    of conveyors.
    """
    return list( bestDirectionForLeastDamage(*conveyor) for conveyor in conveyors )

def getConveyorsFromString(input):
    """
    Take an input list of lines and convert it into a list of conveyor lists.
    Assume 4 lines per conveyor in the input:
        north lasers
        conveyor belt
        south lasers
        blank line
    Throw away the blank line in the output. The last blank line is optional in
    the input.
    """
    return list( input[i:i+3] for i in range(0,len(input),4) )

if __name__ == "__main__":
    from sys import stdin
    for result in bestDirectionsForLeastDamage(getConveyorsFromString(stdin.readlines())):
        print result
