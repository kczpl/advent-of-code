# frozen_string_literal: true

input = File.read("input14.txt", chomp: true)
          .split("\n")
          .map { |i| i.split(" -> ").map { |x| x.split(",").map(&:to_i) } }

cave = Array.new(600) { Array.new(800, ".") }
cave[0][500] = "+"

bottom = 0

input.map do |line|
  line.each_with_index do |coords, index|
    next if line[index + 1].nil?

    x1, y1 = coords
    x2, y2 = line[index + 1]

    bottom = y1 if y1 > bottom

    Range.new([y1, y2].min, [y1, y2].max).each { |y| cave[y][x1] = "#" } if x1 == x2
    Range.new([x1, x2].min, [x1, x2].max).each { |x| cave[y1][x] = "#" } if y1 == y2
  end
end

def print_cave(cave, part)
  File.open("cave_#{part}.txt", "w") do |file|
    cave.map { |r| file << "#{r.join}\n" }
  end
end

buffor = 0
done = false

until done
  buffor += 1
  sand_x, sand_y = 500, 0

  loop do
    if sand_y > bottom
      done = true
      break
    elsif cave[sand_y + 1][sand_x] == "."
      sand_y += 1
    elsif cave[sand_y + 1][sand_x - 1] == "."
      sand_y += 1
      sand_x -= 1
    elsif cave[sand_y + 1][sand_x + 1] == "."
      sand_y += 1
      sand_x += 1
    else
      cave[sand_y][sand_x] = "o"
      break
    end
  end
end

p buffor - 1 # 715
# print_cave(cave, 1)

#========================================v2====================================#

bottom += 2
done = false

until done
  buffor += 1
  sand_x, sand_y = 500, 0

  loop do
    if cave[1][500] == "o" && cave[1][499] == "o" && cave[1][501] == "o"
      done = true
      break
    elsif sand_y >= bottom
      cave[sand_y, sand_x] = "o"
      break
    elsif cave[sand_y + 1][sand_x] == "."
      sand_y += 1
    elsif cave[sand_y + 1][sand_x - 1] == "."
      sand_y += 1
      sand_x -= 1
    elsif cave[sand_y + 1][sand_x + 1] == "."
      sand_y += 1
      sand_x += 1
    else
      cave[sand_y][sand_x] = "o"
      break
    end
  end
end

p buffor - 2 # 25248
# print_cave(cave, 2)
