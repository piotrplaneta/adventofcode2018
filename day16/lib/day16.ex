defmodule Day16 do
  use Bitwise

  @moduledoc """
  Solution for Day16.
  """

  @doc """
  Returns amount of exammples that 3 or more ops satisfy
  """
  def behave_like_three_or_more(input) do
    parse_input(input) |> Enum.filter(& matching_ops(&1) >= 3) |> length()
  end

  @doc """

  Returns given the registers state before and after, and an operation, how many operations can produce that result.

  ## Examples

      iex> Day16.matching_ops({[3, 2, 1, 1], [3, 2, 2, 1], {9, 2, 1, 2}})
      3

  """
  def matching_ops({r_in, r_out, op}) do
    Enum.filter(ops(), & &1.(r_in, op) == r_out)
    |> length()
  end

  defp parse_input(input) do
    String.split(input, "\n")
    |> Enum.chunk_every(4)
    |> Enum.map(fn [r_in | [op | [r_out | [_]]]] ->
      {
        String.split(r_in, "[") |> Enum.at(1) |> String.trim("]")
        |> String.split(", ") |> Enum.map(& String.to_integer(&1)),

        String.split(r_out, "[") |> Enum.at(1) |> String.trim("]")
        |> String.split(", ") |> Enum.map(& String.to_integer(&1)),

        String.split(op, " ") |> Enum.map(& String.to_integer(&1)) |> List.to_tuple(),
      }
    end)
  end

  defp ops do
    [
      fn(r_in, op) ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) + value_at(r_in, second_arg(op))) end,

      fn(r_in, op) ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) + second_arg(op)) end,

      fn(r_in, op) ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) * value_at(r_in, second_arg(op))) end,

      fn(r_in, op) ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) * second_arg(op)) end,

      fn(r_in, op) ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) &&& value_at(r_in, second_arg(op))) end,

      fn(r_in, op) ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) &&& second_arg(op)) end,

      fn(r_in, op) ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) ||| value_at(r_in, second_arg(op))) end,

      fn(r_in, op) ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) ||| second_arg(op)) end,

      fn(r_in, op) ->
        update_at_index(r_in, destination(op), value_at(r_in, second_arg(op))) end,

      fn(r_in, op) ->
        update_at_index(r_in, destination(op), second_arg(op)) end,

      fn(r_in, op) ->
        if first_arg(op) > value_at(r_in, second_arg(op)) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end,

      fn(r_in, op) ->
        if value_at(r_in, first_arg(op)) > second_arg(op) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end,

      fn(r_in, op) ->
        if value_at(r_in, first_arg(op)) > value_at(r_in, second_arg(op)) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end,

      fn(r_in, op) ->
        if first_arg(op) == value_at(r_in, second_arg(op)) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end,

      fn(r_in, op) ->
        if value_at(r_in, first_arg(op)) == second_arg(op) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end,

      fn(r_in, op) ->
        if value_at(r_in, first_arg(op)) == value_at(r_in, second_arg(op)) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end,
    ]
  end

  defp first_arg(op), do: elem(op, 1)
  defp second_arg(op), do: elem(op, 2)
  defp destination(op), do: elem(op, 3)
  defp value_at(register, number), do: Enum.at(register, number)

  defp update_at_index(register, index, value) do
    Enum.take(register, index) ++ [value] ++ Enum.drop(register, index + 1)
  end
end
