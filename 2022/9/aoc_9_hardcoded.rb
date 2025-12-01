# frozen_string_literal: true

MOVES = {
  "L" => [-1, 0],
  "R" => [1, 0],
  "U" => [0, 1],
  "D" => [0, -1]
}.freeze

class Point
  attr_accessor :x, :y

  def initialize(x:, y:)
    @x = x
    @y = y
  end
end

class Head < Point
end

class Tail < Point
end

#========================================v1====================================#

@head = Head.new(x: 0, y: 0)
@tail = Tail.new(x: 0, y: 0)

@initial_head_x = @head.x
@initial_head_y = @head.y

@tails_positions = [[0, 0]]

input = File.readlines("input9.txt", chomp: true).map(&:split)

def move_head(move)
  @initial_head_x = @head.x
  @initial_head_y = @head.y

  coords = MOVES[move.first].map { |i| i * move.last.to_i }
  @head.x += coords.first
  @head.y += coords.last
end

def left_valid(move)
  (@initial_head_x > @tail.x && move.last.to_i >= 3) ||
    (@initial_head_x == @tail.x && move.last.to_i >= 2) ||
    @initial_head_x < @tail.x
end

def right_valid(move)
  (@initial_head_x < @tail.x && move.last.to_i >= 3) ||
    (@initial_head_x == @tail.x && move.last.to_i >= 2) ||
    @initial_head_x > @tail.x
end

def up_valid(move)
  (@initial_head_y < @tail.y && move.last.to_i >= 3) ||
    (@initial_head_y == @tail.y && move.last.to_i >= 2) ||
    @initial_head_y > @tail.y
end

def down_valid(move)
  (@initial_head_y > @tail.y && move.last.to_i >= 3) ||
    (@initial_head_y == @tail.y && move.last.to_i >= 2) ||
    @initial_head_y < @tail.y
end

def valid(move)
  case move.first
  when "L"
    left_valid(move)
  when "R"
    right_valid(move)
  when "U"
    up_valid(move)
  when "D"
    down_valid(move)
  else
    false
  end
end

input.map do |move|
  move_head(move)
  next unless valid(move)

  case move.first
  when "L"
    i = if @initial_head_x > @tail.x && move.last.to_i >= 3
          move.last.to_i - 2
        elsif @initial_head_x == @tail.x && move.last.to_i >= 2
          move.last.to_i - 1
        elsif @initial_head_x < @tail.x
          move.last.to_i
        end

    @tail.y += if @initial_head_y - 1 == @tail.y
                 1
               elsif @initial_head_y + 1 == @tail.y
                 -1
               else
                 0
               end

    i.times do
      @tail.x -= 1
      @tails_positions << [@tail.x, @tail.y]
    end
  when "R"
    i = if @initial_head_x < @tail.x && move.last.to_i >= 3
          move.last.to_i - 2
        elsif @initial_head_x == @tail.x && move.last.to_i >= 2
          move.last.to_i - 1
        elsif @initial_head_x > @tail.x
          move.last.to_i
        end

    @tail.y += if @initial_head_y - 1 == @tail.y
                 1
               elsif @initial_head_y + 1 == @tail.y
                 -1
               else
                 0
               end

    i.times do
      @tail.x += 1
      @tails_positions << [@tail.x, @tail.y]
    end
  when "U"
    i = if @initial_head_y < @tail.y && move.last.to_i >= 3
          move.last.to_i - 2
        elsif @initial_head_y == @tail.y && move.last.to_i >= 2
          move.last.to_i - 1
        elsif @initial_head_y > @tail.y
          move.last.to_i
        end
    @tail.x += if @initial_head_x - 1 == @tail.x
                 1
               elsif @initial_head_x + 1 == @tail.x
                 -1
               else
                 0
               end
    i.times do
      @tail.y += 1
      @tails_positions << [@tail.x, @tail.y]
    end
  when "D"
    i = if @initial_head_y > @tail.y && move.last.to_i >= 3
          move.last.to_i - 2
        elsif @initial_head_y == @tail.y && move.last.to_i >= 2
          move.last.to_i - 1
        elsif @initial_head_y < @tail.y
          move.last.to_i
        end

    @tail.x += if @initial_head_x - 1 == @tail.x
                 1
               elsif @initial_head_x + 1 == @tail.x
                 -1
               else
                 0
               end

    i.times do
      @tail.y -= 1
      @tails_positions << [@tail.x, @tail.y]
    end
  end
end

p @tails_positions.tally.keys.size # 6197

#========================================v2====================================#
