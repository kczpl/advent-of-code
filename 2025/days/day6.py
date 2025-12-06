from collections import defaultdict

from utils import benchmark, read_puzzle


def exec_operation(op: str, nums: list[int]) -> int:
    if op == "*":
        result = 1
        for n in nums:
            result *= n
        return result
    return sum(nums)


@benchmark
def part1(input: str):
    data = input.split("\n")
    columns = len(data)
    operations = data[-1].split()

    tasks = defaultdict(list)
    for num in range(0, columns - 1):
        row = data[num].split()
        for idx, value in enumerate(row):
            tasks[idx].append(int(value))

    result = 0
    for idx, op in enumerate(operations):
        result += exec_operation(op, tasks[idx])
    return result


def calculate_from_right(collection: list[str]):
    operation = collection[-1].strip()
    rows = collection[:-1]

    width = len(collection[0])
    numbers = []

    for col_idx in range(width - 1, -1, -1):  # step -1, stop -1 exclusive
        digits = []
        for row in rows:
            if col_idx < len(row) and row[col_idx] != " ":
                digits.append(row[col_idx])

        if digits:
            number = int("".join(digits))
            numbers.append(number)

    return exec_operation(operation, numbers)


@benchmark
def part2(input: str):
    data = input.split("\n")
    separator_cols = [
        i for i, col in enumerate(zip(*data)) if all(c == " " for c in col)
    ]
    prev_sep = 0
    collections = []

    for s in separator_cols:
        nums = [d[prev_sep:s] for d in data]
        collections.append(nums)
        prev_sep = s + 1

    if prev_sep < len(data[0]):
        nums = [d[prev_sep:] for d in data]
        collections.append(nums)

    total = 0
    for collection in collections:
        result = calculate_from_right(collection)
        total += result

    return total


def solve():
    input = read_puzzle(6)
    part1(input)  # 4771265398012 (0.654 ms)
    part2(input)  # 10695785245101 (2.031 ms)
