# frozen_string_literal: true

require "json"

def compare(s_left, s_right)
  case [s_left.class, s_right.class]
  when [Integer, Integer]
    if s_left < s_right
      1
    elsif s_right < s_left
      0
    end
  when [Array, Array]
    (0...[s_left.length, s_right.length].min).each do |i|
      result = compare(s_left[i], s_right[i])
      return result unless result.nil?
    end
    if s_left.length < s_right.length
      1
    elsif s_left.length > s_right.length
      0
    end
  else
    if s_left.instance_of?(Integer)
      compare([s_left], s_right)
    elsif s_right.instance_of?(Integer)
      compare(s_left, [s_right])
    end
  end
end

input = File.read("input13.txt")
          .split("\n\n")
          .map { |l| l.split("\n") }
          .map { |set| set.map { |i| JSON.parse(i) } }

filtered_signals = []

input.each do |signals|
  s_left, s_right = *signals
  filtered_signals << compare(s_left, s_right)
end

indices = []
filtered_signals.each_with_index do |value, i|
  case value
  when 1
    indices << (i + 1)
  when 0
    next
  end
end

p indices.sum # 5623

#========================================v2====================================#

packs = input + [[[2]], [[6]]]

index_div2 = 1
index_div6 = 1

dividers = packs.last(2)
packs.each do |pack|
  pack.each do |signal|
    index_div2 += 1  if compare(dividers.first, signal) == 0
    index_div6 += 1  if compare(dividers.last, signal) == 0
  end
end

p index_div2 * index_div6 # 20570
