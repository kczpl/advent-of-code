
New day template:
```python
from utils import benchmark, read_puzzle

@benchmark
def part1(input: str):
    pass

@benchmark
def part2(input: str):
    pass


def solve():
    input = read_puzzle(6)
    part1(input)
    part2(input)
```