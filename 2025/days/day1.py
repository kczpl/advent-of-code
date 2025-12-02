from utils import benchmark, read_puzzle


@benchmark
def part1(input: str):
    moves = [(move[0], int(move[1:])) for move in input]
    position = 50
    counter = 0

    for move in moves:
        direction, steps = move

        if direction == "L":
            position = (position - steps) % 100
        elif direction == "R":
            position = (position + steps) % 100

        if position == 0:
            counter += 1

    return counter


@benchmark
def part2(input: str):
    moves = [(move[0], int(move[1:])) for move in input]
    position = 50
    counter = 0

    for direction, distance in moves:
        if direction == "R":
            counter += (position + distance) // 100
            position = (position + distance) % 100
        else:
            if position == 0:
                counter += distance // 100
            elif distance >= position:
                counter += (distance - position) // 100 + 1
            position = (position - distance) % 100

    return counter


def solve():
    input = read_puzzle(1).split("\n")
    part1(input)
    part2(input)
