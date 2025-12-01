# frozen_string_literal: true

input = File.read('input6.txt').chars

marker = input.shift(4)
buffor = 4

input.each do |c|
  break if marker.uniq.length == 4

  marker.shift
  marker << c
  buffor += 1
end

p buffor # 1287

#========================================v2====================================#

input = File.read('input6.txt').chars

marker2 = input.shift(14)
buffor2 = 14

input.each do |c|
  break if marker2.uniq.length == 14

  marker2.shift
  marker2 << c
  buffor2 += 1
end

p buffor2 # 3716
