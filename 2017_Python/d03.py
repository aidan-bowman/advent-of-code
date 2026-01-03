# so here's the plan
# there's a useful pattern here; the bottom right square is ALWAYS an odd square, going 1^2 3^2 5^2 7^2 and so on
# so we can grab a "base" amount and add/subtract from that
# another useful pattern is that the square is symmetrical (e.g. 10 == 14 == 18)
# we can subtract the odd square at the bottom right and then modulo by the base + 1 to get a position relative to a corner
# note: the base + 1 space is on the corner, so we can just set 0 to the previous corner
# ANOTHER pattern: second midpoint (e.g. ---|--) is always closest, just add diff between rel pos and midpoint

import math

def get_lowest_odd_square(num):
    temp = math.floor(math.sqrt(num))
    if temp & 1 == 0:
        return temp - 1
    else:
        return temp
    
def get_non_square_steps(num, base):
    midpoint = int((base + 1) / 2)
    rel_pos = (num - (base * base)) % (base + 1)
    return abs(rel_pos - midpoint) + int((base + 1) / 2)

def get_steps(num):
    base = get_lowest_odd_square(num)
    if num == (base * base):
        return base - 1
    else:
        return get_non_square_steps(num, base)

# PART 2 notes
# FIRST STEP: look at part 1 code and see how to take number and get position
# SECOND STEP: make recursive function: take target num, current position, dict of previous values
# hint: dict.get((x + 1, y), 0) + dict.get((x, y + 1), 0) + ...
# THIRD STEP: you should be able to solve this

def get_pos(num)
    
inp = int(input("Enter number input: "))
print(f"Part 1 solution: {get_steps(inp)}")
# TODO: part 2
