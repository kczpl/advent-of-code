# frozen_string_literal: true

input = File.readlines('input8.txt')

visible = 0
matrix = []

input.map do |l|
  matrix << l.split[0].chars.map(&:to_i)
end

matrix.each_with_index do |row, index|
  if row == matrix.first || row == matrix.last
    visible += row.length
    next
  end
  row.each_with_index do |tree, column|
    if [0, matrix.rindex { true }].include?(column)
      visible += 1
      next
    end

    heighest_neighbours = [
      matrix[0..(index - 1)].map { |r| r[column] }.max,
      matrix[index][column + 1..].max,
      matrix[index][0..column - 1].max,
      matrix[(index + 1)..].map { |r| r[column] }.max
    ]

    visible += 1 if tree > heighest_neighbours.compact.min
  end
end

p visible # 1794

#========================================v2====================================#
@end = input.length - 1
@t_w_crds = {}

input.each_with_index do |line, y|
  line.chars.each_with_index.map { |tree, x| @t_w_crds[[y, x]] = tree.to_i }
end

def up(up, x, y, h)
  (y - up).zero? || @t_w_crds[[y - up, x]].nil? || @t_w_crds[[y - up, x]] >= h
end

def right(right, x, y, h)
  x + right == @end || @t_w_crds[[y, x + right]].nil? || @t_w_crds[[y, x + right]] >= h
end

def down(down, x, y, h)
  y + down == @end || @t_w_crds[[y + down, x]].nil? || @t_w_crds[[y + down, x]] >= h
end

def left(left, x, y, h)
  (x - left).zero? || @t_w_crds[[y, x - left]].nil? || @t_w_crds[[y, x - left]] >= h
end

scores = @t_w_crds.map do |coords, h|
  y, x = coords
  params = [1, 1, 1, 1] # up, right, down, left

  params[0] += 1 until up(params[0], x, y, h)
  params[1] += 1 until right(params[1], x, y, h)
  params[2] += 1 until down(params[2], x, y, h)
  params[3] += 1 until left(params[3], x, y, h)

  result =  params[0] * params[1] * params[2] * params[3]
  result * 0 if x.zero? || x == @end || y.zero? || y == @end
  result
end

p scores.max # 199272
