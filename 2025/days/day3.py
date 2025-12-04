from utils import benchmark, read_puzzle


def parse_input(input: str):
    return [[int(char) for char in line] for line in input.split("\n")]


def cat(left: int, right: int):
    multiplier = 10
    while multiplier <= right:
        multiplier *= 10
    return left * multiplier + right


def find_biggest_battery(bank: list[int], start: int, end: int) -> tuple[int, int]:
    max_val = bank[start]
    max_idx = start
    for i in range(start + 1, end):
        if bank[i] > max_val:
            max_val = bank[i]
            max_idx = i
    return max_val, max_idx


@benchmark
def part1(input: str):
    banks = parse_input(input)
    joltage = 0

    for bank in banks:
        biggest_battery, idx = find_biggest_battery(bank, 0, len(bank) - 1)
        second_battery, _ = find_biggest_battery(bank, idx + 1, len(bank))
        joltage += cat(biggest_battery, second_battery)

    return joltage


@benchmark
def part2(input: str):
    banks = parse_input(input)
    joltage = 0

    def find_sequence(battery: int, idx: int, bank: list[int]) -> int:
        digits = [battery]
        idx += 1
        for i in range(11):
            remaining_needed = 10 - i
            end_idx = len(bank) - remaining_needed
            max_val, max_pos = find_biggest_battery(bank, idx, end_idx)
            digits.append(max_val)
            idx = max_pos + 1
        return int("".join(map(str, digits)))

    for bank in banks:
        biggest_seq = 0
        for idx, battery in enumerate(bank):
            if idx > len(bank) - 12:
                break
            seq = find_sequence(battery, idx, bank)
            if seq > biggest_seq:
                biggest_seq = seq

        joltage += biggest_seq
    return joltage


def solve():
    input = read_puzzle(3)
    part1(input)  # 17263 (1.564 ms)
    part2(input)  # 170731717900423 (57.925 ms) 
