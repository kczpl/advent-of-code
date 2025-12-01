# frozen_string_literal: true

cypher = []

File.readlines('input20.txt', chomp: true)
    .each_with_index { |n, idx| cypher << [idx, n.to_i] }

last = cypher.size
decryptor = cypher.clone

cypher.each do |code|
  next if code.last.zero?

  decryptor_idx = decryptor.find_index(code)

  move = (decryptor_idx + code.last) % (last - 1)
  move -= 1 if move.zero?

  decryptor.delete_at(decryptor.find_index(code))
  decryptor.insert(move, code.last)
end

start_idx = decryptor.find_index(decryptor.detect { |c| c.is_a?(Array) })
decryptor[start_idx] = 0
zero_idx = decryptor.find_index(0)

result = decryptor[(zero_idx + 1000) % last] +
         decryptor[(zero_idx + 2000) % last] +
         decryptor[(zero_idx + 3000) % last]

p result # 9945

#========================================v2====================================#

cypher = []

File.readlines('input20.txt', chomp: true)
    .each_with_index { |n, idx| cypher << [idx, n.to_i * 811_589_153] }

last = cypher.size
decryptor = cypher.clone

10.times do
  cypher.each do |code|
    next if code.last.zero?

    decryptor_idx = decryptor.find_index(code)

    move = (decryptor_idx + code.last) % (last - 1)
    move -= 1 if move.zero?

    decryptor.delete_at(decryptor.find_index(code))
    decryptor.insert(move, code)
  end
end

start_idx = decryptor.find_index(decryptor.detect { |c| c.last.zero? })
decryptor[start_idx] = 0
zero_idx = decryptor.find_index(0)

result = decryptor[(zero_idx + 1000) % last].last +
         decryptor[(zero_idx + 2000) % last].last +
         decryptor[(zero_idx + 3000) % last].last

p result # 3338877775442

# can be definately optimalized to use other data structure than array
# 32s running both
