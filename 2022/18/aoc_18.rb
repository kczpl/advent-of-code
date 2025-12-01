# frozen_string_literal: true

Cube = Struct.new(:x, :y, :z)

def find_neighbors(cube)
  [
    Cube.new(cube.x - 1, cube.y, cube.z), Cube.new(cube.x + 1, cube.y, cube.z),
    Cube.new(cube.x, cube.y - 1, cube.z), Cube.new(cube.x, cube.y + 1, cube.z),
    Cube.new(cube.x, cube.y, cube.z - 1), Cube.new(cube.x, cube.y, cube.z + 1)
  ]
end

cubes = File.readlines("input18.txt", chomp: true ).map { |l| Cube.new(*l.split(",").map(&:to_i )) }

result = cubes.reduce(0) { |sum, cube|
  diff = 6 - find_neighbors(cube).count { |n| cubes.include?(n) }
  sum + diff
}

p result
#========================================v2====================================#

buffor = 0
ranges = [
  (cubes.map(&:x).min - 1)..(cubes.map(&:x).max + 1),
  (cubes.map(&:y).min - 1)..(cubes.map(&:y).max + 1),
  (cubes.map(&:z).min - 1)..(cubes.map(&:z).max + 1)
]

def select_neighbors(ranges, cubes)
  cubes.select { |c|
    ranges[0].include?(c.x) && ranges[1].include?(c.y) && ranges[2].include?(c.z)
  }
end

# cheated with reddit
fuck = []
neighbors = []

fuck << Cube.new(ranges[0].first, ranges[1].first, ranges[2].first)

loop do
  break if fuck.empty?

  cube = fuck.pop

  select_neighbors(ranges, find_neighbors(cube)).each do |neighbor|
    next if neighbors.include?(neighbor)

    if cubes.include?(neighbor)
      buffor += 1
    else
      fuck << neighbor
      neighbors << neighbor
    end
  end
end

p buffor # 2556
