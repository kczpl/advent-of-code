from utils import benchmark, read_puzzle
import re


def parse_ranges(input: str) -> list[range]:
    return [
        range(int(parts[0]), int(parts[1]) + 1)
        for parts in (item.split("-") for item in input.split(","))
    ]


@benchmark
def part1(input: str):
    ranges = parse_ranges(input)

    def is_valid_id(id: int) -> bool:
        s = str(id)
        return s[: len(s) // 2] == s[len(s) // 2 :]

    return sum(id for r in ranges for id in r if is_valid_id(id))


@benchmark
def part2(input: str):
    ranges = parse_ranges(input)

    # TIL: re.compile() avoids recompiling the pattern for every number checked
    # it compiles regex on every part2 call now
    PATTERN = re.compile(r"(.+)\1+")

    def is_valid_id(id: int) -> bool:
        return bool(PATTERN.fullmatch(str(id)))

    return sum(id for r in ranges for id in r if is_valid_id(id))


def solve():
    input = read_puzzle(2)
    part1(input) # 41294979841 (211.860 ms)
    part2(input) # 66500947346 (387.563 ms)

