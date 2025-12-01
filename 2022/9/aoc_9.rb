# frozen_string_literal: true

require "matrix"

# this solution is not mine. It was found in Reddit.
# It was reused and understood, because of great simplicity and education value it provides.

input = File.readlines("input9.txt", chomp: true).map do |l|
  dir, dist = l.split
  [dir, dist.to_i]
end

DIRS = {
  "U" => Vector[0, 1],
  "D" => Vector[0, -1],
  "R" => Vector[1, 0],
  "L" => Vector[-1, 0]
}.freeze

def move(steps, length)
  tail_positions = {}

  knots = Array.new(length) { Vector[0, 0] }
  tail_positions[knots.last] = true

  steps.each do |dir, mag|
    mag.times do
      knots[0] += DIRS[dir]

      (1...knots.length).each do |i|
        prev = knots[i - 1]
        knot = knots[i]

        diff = (prev - knot)

        next unless diff.magnitude >= 2

        normalized = diff.map { |x| x == 0 ? 0 : x / x.abs }
        knots[i] += normalized
      end
      tail_positions[knots.last] = true
    end
  end

  tail_positions.values.count(true)
end

p move(input, 2)
p move(input, 10)
