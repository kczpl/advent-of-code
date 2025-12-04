from utils import benchmark, read_puzzle


DIRECTIONS = [(-1, -1), (-1, 0), (-1, 1), (0, -1), (0, 1), (1, -1), (1, 0), (1, 1)]


def parse_input(input: str):
    return [list[str](line) for line in input.split("\n")]


def check_neighbours(row: int, column: int, matrix: list):
    counter = 0
    for dr, dc in DIRECTIONS:
        r, c = row + dr, column + dc
        if 0 <= r < len(matrix) and 0 <= c < len(matrix[0]):
            if matrix[r][c] == "@":
                counter += 1

    return counter


@benchmark
def part1(input: str):
    matrix = parse_input(input)

    rolls = 0
    for ridx, row in enumerate(matrix):
        for cidx, _ in enumerate(row):
            if matrix[ridx][cidx] == ".":
                continue

            counter = check_neighbours(ridx, cidx, matrix)
            if counter < 4:
                rolls += 1
            else:
                continue
    return rolls


@benchmark
def part2(input: str):
    matrix = parse_input(input)
    rolls = 0

    while True:
        to_remove = []
        for ridx, row in enumerate[list[str]](matrix):
            for cidx, _ in enumerate[str](row):
                if matrix[ridx][cidx] == ".":
                    continue
                counter = check_neighbours(ridx, cidx, matrix)
                if counter < 4:
                    to_remove.append((ridx, cidx))
                    rolls += 1
                else:
                    continue

        if not to_remove:
            break

        for ridx, cidx in to_remove:
            matrix[ridx][cidx] = "."
    return rolls


def solve():
    input = read_puzzle(4)
    part1(input)  # 1533 (42.989 ms)
    part2(input)  # 9206 (235.735 ms)
