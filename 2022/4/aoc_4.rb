# frozen_string_literal: true

input = File.read('input4.txt').lines.map(&:split).flatten.map do |l|
  l.split(',').map { |c| c.split('-').map(&:to_i) }.flatten
end

buffor1 = 0

input.each do |a|
  buffor1 += 1 if a[0..1] == a.minmax || a[2..3] == a.minmax
end

p buffor1 # 498

#========================================v2====================================#

input = File.read('input4.txt').lines.map(&:split).flatten.map do |l|
  l.split(',').map { |c| c.split('-').map(&:to_i) }.sort.flatten
end

buffor2 = 0
input.each do |a|
  buffor2 += 1 if a[0..1].max >= a[2..3].min
end

p buffor2 # 859
