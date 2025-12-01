# frozen_string_literal: true

Directory = Struct.new(:name, :parent, :sub_dirs, :file_size)
input = File.readlines('input7.txt', chomp: true)
input.shift(1)

root = Directory.new('/', nil, [], 0)
current_dir = root

def total_size(dir)
  dir.sub_dirs.map { |sub| total_size(sub) }.sum + dir.file_size
end

def digger(dir)
  s = total_size(dir)

  @buffor += s if s <= 100_000
  dir.sub_dirs.map { |d| digger(d) }
end

input.map do |line|
  if line.include?('$ ls')
    next
  elsif line.include?('dir ')
    current_dir.sub_dirs << Directory.new(line.split.last, current_dir, [], 0)
  elsif line.include?('..')
    current_dir = current_dir.parent
  elsif line.include?('$ cd')
    current_dir = current_dir.sub_dirs.detect { |d| d.name == line.split.last }
  else
    current_dir.file_size += line.split.first.to_i
  end
end

@buffor = 0
digger(root)

p @buffor # 1844187

#========================================v2====================================#

@enough = []
@root_s = total_size(root)

def digger2(dir)
  s = total_size(dir)

  @enough << s if s >= @root_s - 40_000_000
  dir.sub_dirs.map { |d| digger2(d) }
end

digger2(root)

p @enough.min
