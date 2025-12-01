# frozen_string_literal: true

rucksacks = File.readlines("input3.txt").map(&:split).map do |r|
  r[0].chars.each_slice(r[0].length / 2).map(&:join)
end

points = 0

rucksacks.each do |r|
  r.first.each_char do |l|
    next unless r.last.include?(l)

    points += if l.bytes.first < 97
                l.downcase.bytes.first - 70
              else
                l.bytes.first - 96
              end
    break
  end
end

p points # 7848

#========================================v2====================================#

groups = File.readlines("input3.txt").map(&:split).each_slice(3).map(&:flatten)

group_points = 0

groups.each do |g|
  g.sort_by!(&:length)
  g[0].each_char do |l|
    next unless g[1].include?(l) && g[2].include?(l)

    group_points += if l.bytes.first < 97
                      l.downcase.bytes.first - 70
                    else
                      l.bytes.first - 96
                    end
    break
  end
end

p group_points # 2616
