defmodule Day12Test do
  use ExUnit.Case
  doctest Day12

  test "spread plants" do
    assert Day12.spread_plants(elem(test_input(), 0), elem(test_input(), 1)) =~
             "#....##....#####...#######....#.#..##"
  end

  test "sum of pots" do
    assert Day12.sum_of_pots_numbers_after_spreading(elem(test_input(), 0), elem(test_input(), 1)) ==
             325
  end

  defp test_input() do
    {"#..#.#..##......###...###",
     %{
       "...##" => "#",
       "..#.." => "#",
       ".#..." => "#",
       ".#.#." => "#",
       ".#.##" => "#",
       ".##.." => "#",
       ".####" => "#",
       "#.#.#" => "#",
       "#.###" => "#",
       "##.#." => "#",
       "##.##" => "#",
       "###.." => "#",
       "###.#" => "#",
       "####." => "#"
     }}
  end
end
