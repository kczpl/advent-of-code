# frozen_string_literal: true

data = File.readlines("input1.txt")

sums = []

buffor = []

data.each do |line|
  if line == "\n"
    sums << buffor.sum
    buffor.clear
  else
    buffor << line.strip.to_i
  end
end

puts sums.max                    # 1_1
puts sums.sort.reverse[0..2].sum # 1_2
