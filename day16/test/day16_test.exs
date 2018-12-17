defmodule Day16Test do
  use ExUnit.Case
  doctest Day16

  test "behave_like_three_or_more" do
    assert Day16.behave_like_three_or_more(matching_input()) == 651
  end

  test "match_operations" do
    assert Day16.match_operations(matching_input()) == %{
             0 => 7,
             1 => 3,
             2 => 4,
             3 => 5,
             4 => 10,
             5 => 8,
             6 => 0,
             7 => 13,
             8 => 9,
             9 => 1,
             10 => 15,
             11 => 14,
             12 => 6,
             13 => 12,
             14 => 2,
             15 => 11
           }
  end

  test "register_value_after_ops" do
    assert Day16.register_value_after_ops(matching_input(), program_input(), 0) == 706
  end

  defp matching_input() do
    File.read!("test/matching_input")
  end

  defp program_input() do
    File.read!("test/program_input")
  end
end
