import math
from dataclasses import dataclass
from itertools import combinations

from utils import benchmark, read_puzzle


@dataclass(frozen=True)
class Point:
    x: int
    y: int
    z: int


def find_closest_points(points: list[Point]) -> list[tuple]:
    pairs_with_dist = []
    for p1, p2 in combinations(points, 2):
        d = math.dist((p1.x, p1.y, p1.z), (p2.x, p2.y, p2.z))
        pairs_with_dist.append((d, p1, p2))

    pairs_with_dist.sort(key=lambda x: x[0])
    return [(p1, p2) for d, p1, p2 in pairs_with_dist]


def chain(circuits: list[set], p1: Point, p2: Point) -> None:
    c1_idx = None
    c2_idx = None
    for i, c in enumerate(circuits):
        if p1 in c:
            c1_idx = i
        if p2 in c:
            c2_idx = i

    if c1_idx is None and c2_idx is None:
        circuits.append({p1, p2})
    elif c1_idx is not None and c2_idx is None:
        circuits[c1_idx].add(p2)
    elif c1_idx is None and c2_idx is not None:
        circuits[c2_idx].add(p1)
    elif c1_idx is not None and c2_idx is not None and c1_idx != c2_idx:
        if c1_idx < c2_idx:
            circuits[c1_idx].update(circuits[c2_idx])
            circuits.pop(c2_idx)
        else:
            circuits[c2_idx].update(circuits[c1_idx])
            circuits.pop(c1_idx)


@benchmark
def part1(input: str):
    points = [Point(*map(int, row.split(","))) for row in input.splitlines()]
    circuits: list[set] = []
    pairs = find_closest_points(points)

    num_connections = min(
        1000, len(pairs)
    )  # without this it does not work for test but does for results

    for p1, p2 in pairs[:num_connections]:
        chain(circuits, p1, p2)

    sizes = sorted([len(c) for c in circuits], reverse=True)
    return sizes[0] * sizes[1] * sizes[2]


@benchmark
def part2(input: str):
    points = [Point(*map(int, row.split(","))) for row in input.splitlines()]
    circuits: list[set] = []
    pairs = find_closest_points(points)

    for p1, p2 in pairs:
        chain(circuits, p1, p2)

        if len(circuits) == 1 and len(circuits[0]) == len(points):
            return p1.x * p2.x


def solve():
    input = read_puzzle(8)
    part1(input)  # 47040 (285.904 ms)
    part2(input)  # 4884971896 (325.533 ms)
