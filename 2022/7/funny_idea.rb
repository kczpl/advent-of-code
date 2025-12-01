# rubocop: disable all
input = File.readlines('input7.txt')
input.shift(1)

`mkdir root`
path = 'root/'

input.map do |line|
  if line.include?('$ ls')
    next
  elsif line.include?('dir ')
    `mkdir #{path + line.split[1]}`
  elsif line.include?('..')
    path = path.split('/')[0..-2].join('/') << '/'
  elsif line.include?('$ cd')
    path << "#{line.split[2]}/"
  else
    `touch #{path + line.split(' ').first + '_'}`
  end
end

dirs = {}
files = Dir['root/**/*']

files.map do |x|
  data = x.split('/')
  next unless data.last.include?('_')

  if dirs[data[0..data.length - 2].join('/')]
    dirs[data[0..data.length - 2].join('/')] += data.last.split('_').first.to_i
  else
    dirs[data[0..data.length - 2].join('/')] = data.last.split('_').first.to_i
  end
end
buffor = []
groups = []
dirs.map do |dir|
  path = dir.first
  common = dirs.select do |d|
    d.include?(path)
  end

  groups << { dir.first => common.values } if common.values.sum < 100_000
end

d = groups.map do |k, _v|
  buffor << k.values.first.sum
end

p buffor.sum # does not work

