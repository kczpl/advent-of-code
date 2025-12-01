# frozen_string_literal: true

# LEGEND:
# opponent
# A - Rock  1 points
# B - Paper 2 points
# C - Scissors 3 points

# you
# X - Rock  1 points
# Y - Paper 2 points
# Z - Scissors 3 points

# result: loose 0 pts, draw 3 pts, win  6pts

POINTS = {
  "X" => 1,
  "Y" => 2,
  "Z" => 3
}.freeze

CHEAT_POINTS = {
  "X" => 0,
  "Y" => 3,
  "Z" => 6
}.freeze

WIN_IT = {
  "A" => "Y",
  "B" => "Z",
  "C" => "X"
}.freeze

LOSE_IT = {
  "A" => "Z",
  "B" => "X",
  "C" => "Y"
}.freeze

DRAW_IT = {
  "A" => "X",
  "B" => "Y",
  "C" => "Z"
}.freeze

def cheat(game)
  case CHEAT_POINTS[game.last]
  when 3
    DRAW_IT[game.first]
  when 6
    WIN_IT[game.first]
  when 0
    LOSE_IT[game.first]
  end
end

matches = File.readlines("input2.txt").map(&:split)
score = 0
cheat_score = 0

matches.each do |game|
  score += POINTS[game.last]

  if DRAW_IT[game.first] == game.last
    score += 3
  elsif WIN_IT[game.first] == game.last
    score += 6
  end

  cheat_score += CHEAT_POINTS[game.last]
  cheat_score += POINTS[cheat(game)]
end

p score
p cheat_score
