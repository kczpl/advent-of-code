from math import lcm

with open('data/input8.txt', 'r') as file:
    data = file.read()

def part1(data):
    instructions, nodes = data.strip().split('\n\n')

    network = {}
    for line in nodes.split('\n'):
        node, paths = line.split(' = ')
        l, r  =  paths.strip('()').split(', ')
        network[node] = (l, r)

    current = 'AAA'
    steps = 0

    while current != 'ZZZ':
        direction = instructions[steps % len(instructions)]

        if direction == 'L':
            current = network[current][0]
        elif direction == 'R':
            current = network[current][1]

        steps += 1

    print(steps)

def part2(data):
    instructions, nodes = data.strip().split('\n\n')

    network = {}
    for line in nodes.split('\n'):
        node, paths = line.split(' = ')
        left, right = paths.strip('()').split(', ')
        network[node] = (left, right)

    current_positions = [node for node in network.keys() if node.endswith('A')]
    cycle_length = []

    for start in current_positions:
        current = start
        steps = 0

        while not current.endswith('Z'):
            direction = instructions[steps % len(instructions)]
            if direction == 'L':
                current = network[current][0]
            elif direction == 'R':
                current = network[current][1]
            steps += 1

        cycle_length.append(steps)

    # least common multiple of all cycle lengths
    print(lcm(*cycle_length))


part1(data)
part2(data)