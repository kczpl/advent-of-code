require IEx

defmodule Part1 do
  def solve do
    stream = File.stream!("input1.txt")

    {max, _} =
      Enum.reduce(stream, {0, 0}, fn line, {max, buffor} ->
        if line == "\n" do
          {max, 0}
        else
          num =
            line
            |> String.trim()
            |> parse_integer()

          new_buffor = buffor + num
          {Enum.max([max, new_buffor]), new_buffor}
        end
      end)

    max
  end

  defp parse_integer(line) do
    case Integer.parse(line) do
      {num, _} -> num
      :error -> 0
    end
  end
end

defmodule Part2 do
  def solve do
    stream = File.stream!("input1.txt")

    sums =
      Enum.reduce(stream, {[], 0}, fn line, {sums, buffor} ->
        if line == "\n" do
          {[buffor | sums], 0}
        else
          num =
            line
            |> String.trim()
            |> parse_integer()

          new_buffor = buffor + num
          {sums, new_buffor}
        end
      end)
      |> elem(0)

    sums
    |> Enum.sort(&(&1 > &2))
    |> Enum.take(3)
    |> Enum.sum()
  end

  defp parse_integer(line) do
    case Integer.parse(line) do
      {num, _} -> num
      :error -> 0
    end
  end
end

result1 = Part1.solve()
result2 = Part2.solve()
IO.puts("#{result1}")
IO.puts("#{result2}")
