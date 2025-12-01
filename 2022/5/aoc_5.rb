# frozen_string_literal: true

stacks = {
  0 => [],
  1 => [],
  2 => [],
  3 => [],
  4 => [],
  5 => [],
  6 => [],
  7 => [],
  8 => []
}

def move(stacks, qty, from, to)
  stacks[to - 1].unshift(stacks[from - 1].shift(qty).reverse).flatten!
end

def move_v2(stacks, qty, from, to)
  stacks[to - 1].unshift(stacks[from - 1].shift(qty)).flatten!
end

input = File.read('input5.txt')

input.lines do |l|
  break if l == " 1   2   3   4   5   6   7   8   9 \n"

  l.chars.each_slice(4).map(&:join).map(&:split).each_with_index { |a, i| stacks[i] << a[0] }
end

stacks.map { |_k, v| v.compact! }

input.lines.map do |l|
  next unless l[0] == 'm'

  l = l.split
  move(stacks, l[1].to_i, l[3].to_i, l[5].to_i)
end

puts stacks.map { |_k, v| v.first }.join(',').delete('[],')

# V1 GFTNRBZPF
# V2 VRQWPDSGP
