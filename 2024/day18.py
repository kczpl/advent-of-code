with open("data/input18.txt", "r") as f:
    data = f.read()

def part1(data):
    grid = [["." for _ in range(71)] for _ in range(71)]

    coords = []
    for line in data.strip().split("\n"):
        x, y = map(int, line.split(","))
        coords.append((x, y))

    is_valid = (
        lambda x, y, grid: 0 <= x < len(grid)
        and 0 <= y < len(grid[0])
        and grid[y][x] == "."
    )

    def bfs(grid):
        start = (0, 0)
        end = (70, 70)

        dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        queue = [(start, 0)]  # (position, steps)
        visited = {start}

        while queue:
            (x, y), steps = queue.pop(0)
            if (x, y) == end:
                return steps

            for dx, dy in dirs:
                nx, ny = x + dx, y + dy
                new_pos = (nx, ny)
                if new_pos not in visited and is_valid(nx, ny, grid):
                    visited.add(new_pos)
                    queue.append((new_pos, steps + 1))

    for x, y in coords[:1024]:
        grid[y][x] = "#"

    return bfs(grid)


def part2(data):
    coords = []
    for line in data.strip().split("\n"):
        x, y = map(int, line.split(","))
        coords.append((x, y))

    def has_path(blocked_coords):
        grid = [["." for _ in range(71)] for _ in range(71)]
        for x, y in blocked_coords:
            grid[y][x] = "#"

        start = (0, 0)
        end = (70, 70)
        dirs = [(0, 1), (1, 0), (0, -1), (-1, 0)]
        queue = [(start, 0)]
        visited = {start}

        while queue:
            (x, y), steps = queue.pop(0)
            if (x, y) == end:
                return True

            for dx, dy in dirs:
                nx, ny = x + dx, y + dy
                new_pos = (nx, ny)
                if (new_pos not in visited and
                    0 <= nx < 71 and 0 <= ny < 71 and
                    grid[ny][nx] == "."):
                    visited.add(new_pos)
                    queue.append((new_pos, steps + 1))
        return False

    for i in range(len(coords)):
        if not has_path(coords[:i+1]):
            return f"{coords[i][0]},{coords[i][1]}"


print(part1(data))
print(part2(data))