# frozen_string_literal: true

@monkeys = []
@monkeys_v2 = []

class Monkey
  attr_accessor :name, :items, :operation, :tester, :new_level, :count

  def initialize(name:, items:, operation:, tester:)
    @name = name
    @items = items
    @operation = operation
    @tester = tester
    @new_level = nil
    @count = 0
  end

  def monkey_inspect(monkeys)
    items.each do |i|
      monkey_operate(i)
      monkey_test(monkeys)
    end

    @count += items.length
    @items = []
  end

  def monkey_inspect_v2(monkeys, primes_common)
    items.each do |i|
      monkey_operate_v2(i, primes_common)
      monkey_test_v2(monkeys)
    end

    @count += items.length
    @items = []
  end

  def monkey_operate(item)
    operator = if @operation.last == 'old'
                 item
               else
                 @operation.last.to_i
               end

    @new_level = item.send(@operation.first.to_s, operator)
  end

  def monkey_operate_v2(item, primes_common)
    operator = if @operation.last == 'old'
                 item
               else
                 @operation.last.to_i
               end

    @new_level = item.send(@operation.first.to_s, operator) % primes_common
  end

  def monkey_test(monkeys)
    @new_level =  (@new_level / 3.0).floor

    receiver = if (new_level % tester.first).zero?
                 monkeys.detect { |m| m.name == tester[1] }
               else
                 monkeys.detect { |m| m.name == tester[2] }
               end

    receiver.items.push(@new_level)
  end

  def monkey_test_v2(monkeys)
    receiver = if (new_level % tester.first).zero?
                 monkeys.detect { |m| m.name == tester[1] }
               else
                 monkeys.detect { |m| m.name == tester[2] }
               end

    receiver.items.push(@new_level)
  end
end

input = File.readlines('input11.txt', chomp: true).each_slice(7).to_a

input.map do |data|
  monkey = Monkey.new(
    name: data.first.delete(':').split.last.to_i,
    items: data[1].split(':').last.split(',').map(&:to_i),
    operation: data[2].split('=').last.split[1..], # [operator, number]
    tester: [
      data[3].split(' ').last.to_i, # divisible
      data[4].split(' ').last.to_i, # throw if true
      data[5].split(' ').last.to_i  # throw if false
    ]
  )

  monkey_v2 = Monkey.new(
    name: data.first.delete(':').split.last.to_i,
    items: data[1].split(':').last.split(',').map(&:to_i),
    operation: data[2].split('=').last.split[1..], # [operator, number]
    tester: [
      data[3].split(' ').last.to_i, # divisible
      data[4].split(' ').last.to_i, # throw if true
      data[5].split(' ').last.to_i  # throw if false
    ]
  )

  @monkeys << monkey
  @monkeys_v2 << monkey_v2
end

20.times do
  @monkeys.map do |monkey|
    monkey.monkey_inspect(@monkeys)
  end
end

p @monkeys.map(&:count).sort[-2..-1].inject(&:*) # 72884

#========================================v2====================================#
primes_common = @monkeys_v2.map { |m| m.tester.first }.inject(&:*)

10_000.times do
  @monkeys_v2.map do |monkey|
    monkey.monkey_inspect_v2(@monkeys_v2, primes_common)
  end
end

p @monkeys_v2.map(&:count).sort[-2..-1].inject(&:*) # 15310845153
