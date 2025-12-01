# frozen_string_literal: true

CMDS = {
  'noop' => 1,
  'addx' => 2
}.freeze

signals = [20, 60, 100, 140, 180, 220]
x = [1]

File.readlines('input10.txt', chomp: true).map(&:split).each do |l|
  cmd, value = *l
  prev = x[-1]

  [*1..CMDS[cmd]].map { |c| x << (c == CMDS[cmd] ? prev + value.to_i : prev) }
end

p signals.map { |s| x[s - 1] * s }.sum # 14220

#========================================v2====================================#

x.each_with_index do |v, i|
  [*v - 1..v + 1].include?(i % 40) ? print('#') : print('.')

  print "\n" if i % 40 == 39
end

# ####.###...##..###..#....####.####.#..#.
# ...#.#..#.#..#.#..#.#....#.......#.#..#.
# ..#..#..#.#..#.#..#.#....###....#..#..#.
# .#...###..####.###..#....#.....#...#..#.
# #....#.#..#..#.#.#..#....#....#....#..#.
# ####.#..#.#..#.#..#.####.#....####..##..
