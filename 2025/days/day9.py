from utils import benchmark, read_puzzle, read_test_puzzle


def parse_input(input: str) -> list[tuple]:
    return [(int(a), int(b)) for row in input.split("\n") for a, b in [row.split(",")]]


@benchmark
def part1(input: str):
    tiles_coords = parse_input(input)

    max_area = 0
    for idx, (x1, y1) in enumerate(tiles_coords):
        for x2, y2 in tiles_coords[idx + 1 :]:
            area = (abs(x2 - x1) + 1) * (abs(y2 - y1) + 1)
            max_area = max(max_area, area)

    return max_area


@benchmark
def part2(input: str):
    red_ordered = parse_input(input)
    # red_tiles = set(red_ordered)

    edge_green = set()
    for idx in range(len(red_ordered)):
        x1, y1 = red_ordered[idx]
        x2, y2 = red_ordered[(idx + 1) % len(red_ordered)]

        if x1 == x2:  # same column
            for y in range(min(y1, y2) + 1, max(y1, y2)):
                edge_green.add((x1, y))
        else:  # same row
            for x in range(min(x1, x2) + 1, max(x1, x2)):
                edge_green.add((x, y1))

    # boundary = red_tiles | edge_green

    # all_xs = [x for x, y in boundary]
    # all_ys = [y for x, y in boundary]
    # min_x, max_x = min(all_xs) - 1, max(all_xs) + 1
    # min_y, max_y = min(all_ys) - 1, max(all_ys) + 1

    # # Flood fill from outside to find exterior
    # exterior = set()
    # stack = [(min_x, min_y)]
    # while stack:
    #     x, y = stack.pop()
    #     if (x, y) in exterior or (x, y) in boundary:
    #         continue
    #     if x < min_x or x > max_x or y < min_y or y > max_y:
    #         continue
    #     exterior.add((x, y))
    #     stack.extend([(x+1, y), (x-1, y), (x, y+1), (x, y-1)])

    # # Interior = everything in bounding box that's not exterior and not boundary
    # interior = set()
    # for x in range(min_x, max_x + 1):
    #     for y in range(min_y, max_y + 1):
    #         if (x, y) not in exterior and (x, y) not in boundary:
    #             interior.add((x, y))

    # # All valid tiles (red + green edges + green interior)
    # valid_tiles = red_tiles | edge_green | interior

    # # Check each pair of red tiles
    # max_area = 0
    # red_list = list(red_tiles)
    # for i, (x1, y1) in enumerate(red_list):
    #     for x2, y2 in red_list[i + 1:]:
    #         # Check if rectangle is fully within valid tiles
    #         all_valid = True
    #         for x in range(min(x1, x2), max(x1, x2) + 1):
    #             for y in range(min(y1, y2), max(y1, y2) + 1):
    #                 if (x, y) not in valid_tiles:
    #                     all_valid = False
    #                     break
    #             if not all_valid:
    #                 break

    #         if all_valid:
    #             area = (abs(x2 - x1) + 1) * (abs(y2 - y1) + 1)
    #             max_area = max(max_area, area)

    # return max_area


def solve():
    input = read_test_puzzle(9)
    # input = read_puzzle(9)
    part1(input)
    # part2(input)
