from utils import benchmark, read_puzzle


def parse_input(input: str) -> tuple[list[range], list[int]]:
    data = input.split("\n")
    separator = data.index("")
    ranges = [
        range(int(a), int(b) + 1) for s in data[:separator] for a, b in [s.split("-")]
    ]
    ids = [int(id) for id in data[separator + 1 :]]
    return ranges, ids


@benchmark
def part1(input: str):
    ranges, ids = parse_input(input)

    counter = 0
    for id in ids:
        for r in ranges:
            if id in r:
                counter += 1
                break
    return counter


@benchmark
def part2(input: str):
    ranges, ids = parse_input(input)

    def get_uncovered(r: range, covered_ranges: list[range]) -> list[range]:
        uncovered = [r]

        for covered in covered_ranges:
            new_uncovered = []
            for ur in uncovered:
                # no overlap
                if ur.stop <= covered.start or covered.stop <= ur.start:
                    new_uncovered.append(ur)
                else:
                    # left part
                    if ur.start < covered.start:
                        new_uncovered.append(range(ur.start, covered.start))
                    # right part
                    if ur.stop > covered.stop:
                        new_uncovered.append(range(covered.stop, ur.stop))
            uncovered = new_uncovered

        return uncovered

    counter = 0
    covered_ranges = []
    for r in ranges:
        uncovered = get_uncovered(r, covered_ranges)
        counter += sum(len(u) for u in uncovered)
        covered_ranges.append(r)

    return counter


def solve():
    input = read_puzzle(5)
    part1(input) # 733 (3.601 ms)
    part2(input) # 345821388687084 (1.242 ms)
