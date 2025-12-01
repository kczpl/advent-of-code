# frozen_string_literal: true
# https://adventofcode.com/2021/day/6

input = File.read('input.txt').split(',').map(&:to_i)

80.times do
  buffor = 0

  input = input.map do |fish|
    if fish - 1 == -1
      fish = 6
      buffor += 1
    else
      fish -= 1
    end
    fish
  end

  buffor.times do
    input << 8
  end
end

p input.length # 386755

#==============================v2===============================================

input = File.read('input.txt').split(',').map(&:to_i)

fishes = [*0..8].map { |i| input.count(i) }

256.times do
  breed = fishes[0]

  [*0..7].map do |i|
    fishes[i] = fishes[i + 1]
  end

  fishes[6] += breed
  fishes[8] = breed
end

p fishes.sum # 1732731810807

#==============================v2===============================================

input = File.read('input.txt').split(',').map(&:to_i)

fishes = [*0..8].map { |i| input.count(i) }

256.times do
  buffor = fishes[0]

  fishes[0] = fishes[1]
  fishes[1] = fishes[2]
  fishes[2] = fishes[3]
  fishes[3] = fishes[4]
  fishes[4] = fishes[5]
  fishes[5] = fishes[6]
  fishes[6] = fishes[7] + buffor
  fishes[7] = fishes[8]
  fishes[8] = buffor
end

p fishes.sum # 1732731810807
