from functools import cache

from utils import benchmark, read_puzzle, read_test_puzzle


def parse_input(input: str):
    return [list(row) for row in input.split("\n")]


@benchmark
def part1(input: str):
    matrix = parse_input(input)
    beams_idxs = {matrix[0].index("S")}
    counter = 0
    for idx, row in enumerate(matrix):
        if idx == 0:
            continue

        new_beams = set()
        for bidx in beams_idxs:
            if matrix[idx][bidx] == "^":
                new_beams.add(bidx - 1)
                new_beams.add(bidx + 1)
                counter += 1
            else:
                new_beams.add(bidx)
        beams_idxs = new_beams

    return counter


@benchmark
def part2(input: str):
    matrix = parse_input(input)

    start = matrix[0].index("S")
    rows_num = len(matrix)

    splitters = {
        (ridx, cidx)
        for ridx, row in enumerate(matrix)
        for cidx, c in enumerate(row)
        if c == "^"
    }

    @cache
    def traverse(col, row):
        while row < rows_num:
            if (row, col) in splitters:
                left_count = traverse(col - 1, row)
                right_count = traverse(col + 1, row)
                return left_count + right_count
            row += 1
        return 1

    return traverse(start, 0)


def solve():
    input = read_puzzle(7)
    part1(input)  # 1678 (0.379 ms)
    part2(input)  # 357525737893560 (1.332 ms)
