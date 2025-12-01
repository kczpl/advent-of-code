# frozen_string_literal: true

require "pry"
Coords = Struct.new(:x, :y)

# [ sensor[x,y] ,  beacon[x,y]  ]
input = File.read("input15.txt").split("\n")
          .map { |l| l.split(/[xy]=(-?\d+)/).map(&:to_i) }
          .map { |sl| [Coords.new(sl[1], sl[3]), Coords.new(sl[-3], sl[-1])] }

@banned_x = []

def calculate_distance(sensor, beacon)
  (beacon.x - sensor.x).abs + (beacon.y - sensor.y).abs
end

def to_row_dist(target_y, sensor, dist)
  y_dist = (sensor.y - target_y).abs

  return if y_dist > dist

  range = (y_dist.abs - dist.abs)

  [*0..range.abs].each do |i|
    @banned_x << (sensor.x + i)
    @banned_x << (sensor.x - i)
  end
end

target_y = 2_000_000
beacons_in_target = []
input.each { |pair| pair.last.y == target_y ? beacons_in_target << pair.last.x : nil }

input.map do |sensor_beacon|
  sensor, beacon = sensor_beacon

  dist = calculate_distance(sensor, beacon)
  to_row_dist(target_y, sensor, dist)
end

p @banned_x.uniq.count - beacons_in_target.uniq.count # 5040643

#========================================v2====================================#
RANGE = 4_000_000

sensors = input.map do |sensor_beacon|
  sensor, beacon = sensor_beacon

  [sensor, calculate_distance(sensor, beacon)]
end

# 4 lines y=ax +b
# all possible y-intercept for sensor

lines = sensors.map { |sens, dist|
  [
    [sens.x - sens.y + dist, sens.x - sens.y - dist], # a = 1
    [sens.x + sens.y + dist, sens.x + sens.y - dist]  # a = -1
  ]
}.transpose.map(&:flatten)

lines[0].product( lines[1] )
  .map { |c1, c2| (c2 - c1) / 2 }
  .select { |y| (0..RANGE ).include?(y) }
  .uniq
  .each do |row|
  ranges = sensors.map { |sens, dist|
    dist -= (row - sens.y ).abs
    [sens.x - dist, sens.x + dist]
  }.select { |a, b| b > a }.sort

  col = ranges.first.first
  ranges.each do |x1, x2|
    if col >= x1 && col <= x2
      col = x2 + 1
    elsif x1 > col
      return p row + (col * 4_000_000) # 11016575214126
    end
  end
end
