from dataclasses import dataclass, field
# from functools import lru_cache

from utils import benchmark, read_puzzle, read_test_puzzle


@dataclass
class Node:
    name: str
    outputs: list = field(default_factory=list)


def parse_input(input: str):
    nodes = {}
    for row in input.split("\n"):
        name, targets = row.split(": ")
        if name not in nodes:
            nodes[name] = Node(name)
        for target in targets.split():
            nodes[name].outputs.append(target)
    return nodes


def find_all_paths(nodes, start, end):
    paths = []
    stack = [(start, [start])]

    while stack:
        node, path = stack.pop()

        if node == end:
            paths.append(path)
            continue

        if node not in nodes:
            continue

        for next_node in nodes[node].outputs:
            stack.append((next_node, path + [next_node]))
    return paths


# def count_paths(nodes, start, end, required):
#     graph = {name: node.outputs for name, node in nodes.items()}
#     required = frozenset(required)

#     @lru_cache(maxsize=None)
#     def dfs(node, seen):
#         if node == end:
#             return 1 if seen == required else 0
#         if node not in graph:
#             return 0
#         if node in required:
#             seen = seen | {node}
#         return sum(dfs(n, seen) for n in graph[node])

#     return dfs(start, frozenset())


@benchmark
def part1(input: str):
    nodes = parse_input(input)
    paths = find_all_paths(nodes, "you", "out")
    return len(paths)


@benchmark
def part2(input: str):
    pass
    # nodes = parse_input(input)
    # return count_paths(nodes, "svr", "out", {"dac", "fft"})


def solve():
    # input = read_test_puzzle(11)
    input = read_puzzle(11)
    part1(input)
    part2(input)
