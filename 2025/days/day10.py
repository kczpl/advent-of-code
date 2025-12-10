from itertools import product

from utils import benchmark, read_puzzle


def parse_input(input: str):
    lines = []
    for row in input.split("\n"):
        parts = row.split()

        lights = None
        buttons = []
        joltage = []

        for part in parts:
            if part.startswith("["):
                parsed_lights = part.strip("[]")
                lights = [1 if char == "#" else 0 for char in parsed_lights]
            elif part.startswith("("):
                buttons.append(tuple(int(x) for x in part.strip("()").split(",")))
            elif part.startswith("{"):
                joltage = [int(x) for x in part.strip("{}").split(",")]

        lines.append({"lights": lights, "buttons": buttons, "joltage": joltage})
    return lines


def find_fewest_press(target: list, buttons: list[tuple]):
    num_lights = len(target)

    valid_presses = []

    for presses in product([0, 1], repeat=len(buttons)):
        state = [0] * num_lights

        for btn_idx, pressed in enumerate(presses):
            if pressed == 1:
                for light_idx in buttons[btn_idx]:
                    state[light_idx] ^= 1  # toggle XOR

        if state == target:
            valid_presses.append(sum(presses))

    return min(valid_presses)


def find_fewest_press_with_joltage(target: list[int], buttons: list[tuple]):
    num_counters = len(target)
    max_presses = max(target) + 1

    valid_presses = []

    for presses in product(range(max_presses), repeat=len(buttons)):
        state = [0] * num_counters

        for btn_idx, press_count in enumerate(presses):
            for counter_idx in buttons[btn_idx]:
                state[counter_idx] += press_count

        if state == target:
            valid_presses.append(sum(presses))

    return min(valid_presses)


@benchmark
def part1(input: str):
    lines = parse_input(input)

    sum = 0
    for line in lines:
        sum += find_fewest_press(line["lights"], line["buttons"])
    return sum


@benchmark
def part2(input: str):
    lines = parse_input(input)
    total = 0
    for idx, line in enumerate(lines):
        # print(f"""line idx: {idx}""")
        total += find_fewest_press_with_joltage(line["joltage"], line["buttons"])
    return total


def solve():
    # input = read_test_puzzle(10)
    input = read_puzzle(10)
    part1(input)
    # part2(input) # brute force does not works
