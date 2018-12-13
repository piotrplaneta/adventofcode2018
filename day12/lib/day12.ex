defmodule Day12 do
  @moduledoc """
  Documentation for Day12.
  """

  @doc """
  Returns sum of pots including '#' after guven number of steps
  """
  def sum_of_pots_numbers_after_spreading(initial_state, rules, steps \\ 20) do
    spread_plants(initial_state, rules, steps)
    |> String.codepoints()
    |> Enum.with_index()
    |> Enum.reduce(0, fn {value, index}, sum ->
      if value == "#" do
        sum + (index - steps * 3)
      else
        sum
      end
    end)
  end

  @doc """
  Spread plants and returns how they look after given amount of steps.
  """
  def spread_plants(initial_state, rules, steps \\ 20) do
    Enum.reduce(1..steps, initial_state, fn _, plants ->
      Enum.reduce(generate_extended_state_chunks(plants), "", fn step, acc ->
        step_string = List.to_string(step)

        if Map.has_key?(rules, step_string) do
          acc <> rules[step_string]
        else
          acc <> "."
        end
      end)
    end)
  end

  defp generate_extended_state_chunks(initial_state) do
    ("....." <> initial_state <> "......")
    |> String.codepoints()
    |> Enum.chunk_every(5, 1)
  end
end
