# frozen_string_literal: true

graph = []

Coords = Struct.new(:x, :y)
GraphNode = Struct.new(:id, :coords, :value, :neighbor_ids)

matrix = File.readlines('input12.txt', chomp: true).map do |row|
  row.chars.map do |c|
    case c
    when 'S'
      0
    when 'E'
      123 - 96
    else
      c.bytes.first - 96
    end
  end
end

buffor = 0

matrix.each_with_index do |row, row_index|
  row.each_with_index do |value, column|
    graph << GraphNode.new(buffor, Coords.new(column, row_index), value)
    buffor += 1
  end
end

graph.map do |node|
  starter = false
  ender = false

  if node.value == 0
    starter = true
    node.value + 1
  end

  if node.value == 27
    ender = true
    node.value - 1
  end

  # can be optimalize via using Hash instead of Array
  up = graph.find do |n|
    n.coords.y == node.coords.y - 1 && n.coords.x == node.coords.x && (n.value <= node.value + 1)
  end
  right = graph.find do |n|
    n.coords.y == node.coords.y && n.coords.x == node.coords.x + 1 && (n.value <= node.value + 1)
  end
  left = graph.find do |n|
    n.coords.y == node.coords.y && n.coords.x == node.coords.x - 1 && (n.value <= node.value + 1)
  end
  down = graph.find do |n|
    n.coords.y == node.coords.y + 1 && n.coords.x == node.coords.x && (n.value <= node.value + 1)
  end

  node.neighbor_ids = [up, right, left, down].compact.map(&:id)

  node.value - 1 if starter
  node.value + 1 if ender
end

def path(prev, end_id)
  path = []

  path << end_id

  while prev[end_id]
    end_id = prev[end_id]
    path << end_id
  end

  path
end

def bfs(graph, start_id, end_id)
  queue = []
  queue << start_id

  visited = { start_id => 1 }
  prev = {}

  until queue.empty?
    current_id = queue.shift
    return path(prev, end_id) if current_id == end_id

    graph[current_id].each do |neighbor_id|
      next if visited[neighbor_id]

      queue << neighbor_id
      prev[neighbor_id] = current_id

      visited[neighbor_id] = 1
    end
  end
end

start_node = graph.detect { |n| n.value.zero? }.id
end_node = graph.detect { |n| n.value == 27 }.id

hash_graph = {}
graph.map { |n| hash_graph[n.id] = n.neighbor_ids }

visited = bfs(hash_graph, start_node, end_node)
p visited.count - 1 # 534

#========================================v2====================================#

lowest = graph.select { |n| n.value == 1 }

min = Float::INFINITY

lowest.map do |s|
  visited2 = bfs(hash_graph, s.id, end_node)

  min = visited2.count if visited2 && visited2.count < min
end

p min - 1 # 525
