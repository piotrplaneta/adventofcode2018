defmodule Day15Test do
  use ExUnit.Case
  doctest Day15

  test "movement tick game_state" do
    {_, _, initial_state} = Day15.parse_input(movement_test_input())
    game_state = Day15.tick(initial_state)

    assert game_state[{4, 3}].type == :elf
  end

  test "movement tick gremlin" do
    {_, _, initial_state} = Day15.parse_input(movement_test_input())
    game_state = Day15.tick(initial_state)

    assert game_state[{2, 1}].type == :gremlin
  end

  test "movement tick gremlin with obstacle" do
    {_, _, initial_state} = Day15.parse_input(movement_test_input())
    game_state = Day15.tick(initial_state)

    assert game_state[{3, 7}].type == :gremlin
  end

  test "goblins score" do
    {_, _, initial_state} = Day15.parse_input(tick_test_input())
    {:gremlins, score} = Day15.winning_score(initial_state)

    assert score == 27730
  end

  defp movement_test_input() do
    File.read!("test/movement_test_input")
    |> String.split("\n")
  end

  defp tick_test_input() do
    File.read!("test/tick_test_input")
    |> String.split("\n")
  end
end
