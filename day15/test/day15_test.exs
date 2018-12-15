defmodule Day15Test do
  use ExUnit.Case
  doctest Day15

  test "movement tick elf" do
    parsed_input = Day15.parse_input(movement_test_input())
    {elfs, _, _} = Day15.movement_tick(parsed_input)

    assert Enum.at(elfs, 0).position == {4, 3}
  end

  test "movement tick game_state" do
    parsed_input = Day15.parse_input(movement_test_input())
    {_, _, game_state} = Day15.movement_tick(parsed_input)

    assert game_state[{4, 3}].type == :elf
  end

  test "movement tick gremlin" do
    parsed_input = Day15.parse_input(movement_test_input())
    {_, gremlins, _} = Day15.movement_tick(parsed_input)

    assert Enum.at(gremlins, 0).position == {2, 1}
  end

  test "movement tick gremlin with obstacle" do
    parsed_input = Day15.parse_input(movement_test_input())
    {_, gremlins, _} = Day15.movement_tick(parsed_input)

    assert Enum.at(gremlins, -2).position == {3, 7}
  end

  defp movement_test_input() do
    File.read!("test/movement_test_input")
    |> String.split("\n")
  end
end
