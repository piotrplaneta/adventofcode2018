defmodule Day16Test do
  use ExUnit.Case
  doctest Day16

  test "behave_like_three_or_more" do
    assert Day16.behave_like_three_or_more(input()) == 651
  end

  test "match_operations" do
    assert Day16.match_operations(input()) == 651
  end

  defp input() do
    File.read!("test/input")
  end
end
