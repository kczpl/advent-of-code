require IEx

defmodule AdventOfCode.Day3 do
  def part1 do
    data = read_input("input3.txt")
  end

  def part2 do
  end

  def read_input(file_path) do
    case File.read(file_path) do
      {:ok, data} ->
        String.split(data, "\n", trim: true)

      {:error, reason} ->
        IO.puts("Error reading file: #{reason}")
        []
    end
  end
end

IO.puts("Part 1 Points: #{AdventOfCode.Day3.part1()}")
IO.puts("Part 2 Points: #{AdventOfCode.Day3.part2()}")
