from functools import cache

with open("data/input19.txt") as f:
    data = f.read()

@cache
def valid(patterns, design):
    # if design is empty there is a match
    if not design:
        return True


    for pattern in patterns:
        if design.startswith(pattern):
            remaining = design[len(pattern):]
            if valid(patterns, remaining):
                return True

    # if no match, return false
    return False

def part1(data):
    patterns, designs = data.strip().split("\n\n")
    patterns = tuple(p.strip() for p in patterns.split(','))
    sum = 0

    for design in designs.splitlines():
        if valid(patterns, design):
            sum += 1
    print(sum)

@cache
def count_ways(patterns, design):
    # if empty, all patterns checked and one valid way
    if not design:
        return 1

    total = 0
    for pattern in patterns:
        if design.startswith(pattern):
            remaining = design[len(pattern):]
            ways = count_ways(patterns, remaining)
            total += ways

    return total

def part2(data):
    patterns, designs = data.strip().split("\n\n")
    patterns = tuple(p.strip() for p in patterns.split(','))

    total = 0
    for design in designs.splitlines():
        ways = count_ways(patterns, design)
        total += ways
    print(total)

part1(data)
part2(data)