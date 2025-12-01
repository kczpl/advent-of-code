require IEx

defmodule Day5.Common do
  def read_input(file_path) do
    case File.read(file_path) do
      {:ok, data} ->
        String.split(data, "\n\n", trim: true)

      {:error, reason} ->
        IO.puts("Error reading file: #{reason}")
        []
    end
  end
end

defmodule Day5.Part1 do
  def solve(data) do
    init_seeds = parse_init_seeds(data)
    data = Enum.drop(data, 1)

    almanac = parse_almanac(data)

    Enum.map(init_seeds, fn seed ->
      Enum.reduce(almanac, seed, fn maps, value ->
        Enum.reduce_while(maps, value, fn [dest, source, range], acc ->
          if acc in source..(source + range - 1) do
            {:halt, dest + (acc - source)}
          else
            {:cont, acc}
          end
        end)
      end)
    end)
    |> Enum.min()
  end

  def parse_init_seeds(data) do
    data
    |> List.first()
    |> String.split(": ")
    |> Enum.drop(1)
    |> hd()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
  end

  def parse_almanac(data) do
    data
    |> Enum.map(fn x ->
      x
      |> String.split("\n")
      |> Enum.drop(1)
      |> Enum.reject(&(&1 == ""))
      |> Enum.map(fn set ->
        set
        |> String.split(" ")
        |> Enum.map(&String.to_integer/1)
      end)
    end)
  end
end

defmodule Day5.Part2 do
  def solve(data) do
    init_seeds = parse_init_seeds(data)

    almanac = parse_almanac(data)
    almanac = Enum.drop(almanac, 1)

    almanac =
      Enum.map(almanac, fn x ->
        Enum.map(x, fn [dest, source, range] ->
          [dest..(dest + range - 1), source..(source + range - 1), dest, source, range]
        end)
      end)

    val =
      Enum.map(init_seeds, fn seed_range ->
        Enum.map(almanac, fn map ->
          find_localization(seed_range, map)
        end)

        # |> Enum.min()
      end)

    # brute force ftw
    Enum.min(val)
  end

  def find_localization(seed_range, map) do
    Enum.reduce(map, seed_range, fn maps, s_range ->
      Enum.reduce_while(maps, s_range, fn map, s_range ->
        if intersection(s_range, hd(map)) do
          a = intersection(s_range, hd(map))

          #  b=Enum.reduce_while(map, value, fn [dest, source, range], acc ->
          #   if acc in source..(source + range - 1) do
          #     {:halt, dest + (acc - source)}
          #   else
          #     {:cont, acc}
          #   end
          # end)
          IEx.pry()

          {:halt, a}
        else
          {:cont, s_range}
        end
      end)
    end)
  end

  def intersection(range1, range2) do
    start = max(Enum.min(range1), Enum.min(range2))
    finish = min(Enum.max(range1), Enum.max(range2))

    if start <= finish do
      start..finish
    else
      nil
    end
  end

  def parse_init_seeds(data) do
    data
    |> List.first()
    |> String.split(": ")
    |> Enum.drop(1)
    |> hd()
    |> String.split(" ")
    |> Enum.map(&String.to_integer/1)
    |> Enum.chunk_every(2)
    |> Enum.map(fn [start, length] -> start..(start + length - 1) end)
  end

  def parse_almanac(data) do
    data
    |> Enum.map(fn x ->
      x
      |> String.split("\n")
      |> Enum.drop(1)
      |> Enum.reject(&(&1 == ""))
      |> Enum.map(fn set ->
        set
        |> String.split(" ")
        |> Enum.map(&String.to_integer/1)
      end)
    end)
  end
end

data = Day5.Common.read_input("input_5.txt")
# => 88151870
# IO.puts("Part 1: #{Day5.Part1.solve(data)}")
# =>
IO.puts("Part 2 : #{Day5.Part2.solve(data)}")
