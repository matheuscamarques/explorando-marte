defmodule Commandline.CLI do
  def main(_) do
    array_length = IO.read(:stdio, :line)
    array = IO.read(:stdio, :line)

    array_length
        |> String.trim
        |> String.to_integer
        |> IO.puts

    array
        |> String.split(" ")
        |> Enum.map(fn n -> String.to_integer(n) end)
        |> IO.inspect
  end

end
