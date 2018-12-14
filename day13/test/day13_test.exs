defmodule Day13Test do
  use ExUnit.Case
  doctest Day13

  test "first crash" do
    assert Day13.first_crash(test_input()) == {7, 3}
  end

  test "last car standing" do
    assert Day13.last_car_standing(last_car_standing_test_input()) == {6, 4}
  end

  defp test_input() do
    File.read!("test/input")
    |> String.split("\n")
  end

  defp last_car_standing_test_input() do
    File.read!("test/last_car_standing_input")
    |> String.split("\n")
  end
end
