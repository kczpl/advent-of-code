# frozen_string_literal: true

require 'pry'

@monkeys = {}

File.readlines('input21.txt', chomp: true).map do |l|
  name, value = l.split(':')

  @monkeys[name] = if value.scan(/\d+/).any?
                     value.strip.to_i
                   else
                     value.strip
                   end
end

def yield_number(monkey)
  return monkey if monkey.is_a?(Integer)

  value = monkey.split
  v1 = yield_number(@monkeys[value.first])
  v2 = yield_number(@monkeys[value.last])

  v1.send(value[1].to_sym, v2)
end

result = yield_number(@monkeys['root'])

p result # 152479825094094

#========================================v2====================================#

def find_number(monkey)
  return monkey if monkey.is_a?(Integer)

  value = monkey.split
  v1 = find_number(@monkeys[value.first])
  v2 = find_number(@monkeys[value.last])

  v1.send(value[1].to_sym, v2)
end

def find_root(monkey)
  value = monkey.split
  v1 = find_number(@monkeys[value.first])
  v2 = find_number(@monkeys[value.last])

  [v1, v2]
end

# brute force bitch
# printing and 'binary search' when  root.first < root.last for 33_360_562_000_000
@buffor = 3_360_561_000_000

loop do
  @buffor += 1

  @monkeys['humn'] = @buffor

  root = find_root(@monkeys['root'])

  break if root.first == root.last
end

p @buffor # 3360561285172
