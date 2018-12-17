defmodule Day16 do
  use Bitwise

  def behave_like_three_or_more(input) do
    parse_matching_input(input) |> Enum.filter(&(matching_ops(&1) >= 3)) |> length()
  end

  def register_value_after_ops(operations_input, program_input, register_index) do
    operations = match_operations(operations_input)

    program = parse_program_input(program_input)

    registers =
      Enum.reduce(program, [0, 0, 0, 0], fn op, current_registers ->
        Enum.at(ops(), operations[op_number(op)]).(current_registers, op)
      end)

    Enum.at(registers, register_index)
  end

  def match_operations(input) do
    matching_operations =
      parse_matching_input(input)
      |> Enum.map(fn {r_in, r_out, op} ->
        Enum.with_index(ops())
        |> Enum.filter(&(elem(&1, 0).(r_in, op) == r_out))
        |> Enum.map(&{op_number(op), elem(&1, 1)})
        |> Enum.group_by(&elem(&1, 0), &elem(&1, 1))
      end)
      |> Enum.map(&Map.to_list(&1))
      |> Enum.map(&Enum.at(&1, 0))
      |> Enum.group_by(&elem(&1, 0), &elem(&1, 1))
      |> Enum.map(fn {op_number, possible_operations} ->
        {op_number, Enum.map(possible_operations, &MapSet.new(&1))}
      end)
      |> Enum.map(fn {op_number, possible_operations} ->
        {op_number, Enum.reduce(possible_operations, &MapSet.intersection/2)}
      end)
      |> Enum.map(fn {op_number, possible_operations} ->
        {op_number, MapSet.to_list(possible_operations)}
      end)

    match_one_to_one(matching_operations, [])
  end

  defp match_one_to_one(_, already_matched) when length(already_matched) == 16 do
    Enum.into(already_matched, %{})
  end

  defp match_one_to_one(possible_matches, already_matched) do
    new_matched =
      Enum.filter(possible_matches, fn {_, possible_ops} -> length(possible_ops) == 1 end)
      |> Enum.map(fn {k, v} -> {k, Enum.at(v, 0)} end)

    new_possible_matches =
      Enum.map(possible_matches, fn {op_num, possible_ops} ->
        {op_num,
         Enum.reject(possible_ops, fn possible_op ->
           Enum.any?(new_matched, &(elem(&1, 1) == possible_op))
         end)}
      end)

    match_one_to_one(new_possible_matches, already_matched ++ new_matched)
  end

  defp matching_ops({r_in, r_out, op}) do
    Enum.filter(ops(), &(&1.(r_in, op) == r_out))
    |> length()
  end

  defp parse_program_input(input) do
    String.split(input, "\n")
    |> Enum.map(fn line ->
      String.split(line, " ") |> Enum.map(&String.to_integer(&1)) |> List.to_tuple()
    end)
  end

  defp parse_matching_input(input) do
    String.split(input, "\n")
    |> Enum.chunk_every(4)
    |> Enum.map(fn [r_in | [op | [r_out | [_]]]] ->
      {
        String.split(r_in, "[")
        |> Enum.at(1)
        |> String.trim("]")
        |> String.split(", ")
        |> Enum.map(&String.to_integer(&1)),
        String.split(r_out, "[")
        |> Enum.at(1)
        |> String.trim("]")
        |> String.split(", ")
        |> Enum.map(&String.to_integer(&1)),
        String.split(op, " ") |> Enum.map(&String.to_integer(&1)) |> List.to_tuple()
      }
    end)
  end

  defp ops do
    [
      fn r_in, op ->
        update_at_index(
          r_in,
          destination(op),
          value_at(r_in, first_arg(op)) + value_at(r_in, second_arg(op))
        )
      end,
      fn r_in, op ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) + second_arg(op))
      end,
      fn r_in, op ->
        update_at_index(
          r_in,
          destination(op),
          value_at(r_in, first_arg(op)) * value_at(r_in, second_arg(op))
        )
      end,
      fn r_in, op ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) * second_arg(op))
      end,
      fn r_in, op ->
        update_at_index(
          r_in,
          destination(op),
          value_at(r_in, first_arg(op)) &&& value_at(r_in, second_arg(op))
        )
      end,
      fn r_in, op ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) &&& second_arg(op))
      end,
      fn r_in, op ->
        update_at_index(
          r_in,
          destination(op),
          value_at(r_in, first_arg(op)) ||| value_at(r_in, second_arg(op))
        )
      end,
      fn r_in, op ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)) ||| second_arg(op))
      end,
      fn r_in, op ->
        update_at_index(r_in, destination(op), value_at(r_in, first_arg(op)))
      end,
      fn r_in, op ->
        update_at_index(r_in, destination(op), first_arg(op))
      end,
      fn r_in, op ->
        if first_arg(op) > value_at(r_in, second_arg(op)) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end,
      fn r_in, op ->
        if value_at(r_in, first_arg(op)) > second_arg(op) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end,
      fn r_in, op ->
        if value_at(r_in, first_arg(op)) > value_at(r_in, second_arg(op)) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end,
      fn r_in, op ->
        if first_arg(op) == value_at(r_in, second_arg(op)) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end,
      fn r_in, op ->
        if value_at(r_in, first_arg(op)) == second_arg(op) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end,
      fn r_in, op ->
        if value_at(r_in, first_arg(op)) == value_at(r_in, second_arg(op)) do
          update_at_index(r_in, destination(op), 1)
        else
          update_at_index(r_in, destination(op), 0)
        end
      end
    ]
  end

  defp op_number(op), do: elem(op, 0)
  defp first_arg(op), do: elem(op, 1)
  defp second_arg(op), do: elem(op, 2)
  defp destination(op), do: elem(op, 3)
  defp value_at(register, number), do: Enum.at(register, number)

  defp update_at_index(register, index, value) do
    Enum.take(register, index) ++ [value] ++ Enum.drop(register, index + 1)
  end
end
