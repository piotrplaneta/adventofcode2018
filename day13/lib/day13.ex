defmodule Day13 do
  defmodule Car do
    defstruct position: {}, direction: {}, next_turn: :left

    def move(car, track_piece) when track_piece in ["-", "|"] do
      new_position =
        {elem(car.position, 0) + elem(car.direction, 0),
         elem(car.position, 1) + elem(car.direction, 1)}

      %Car{position: new_position, direction: car.direction, next_turn: car.next_turn}
    end

    def move(car, "/") do
      new_position =
        {elem(car.position, 0) - elem(car.direction, 1),
         elem(car.position, 1) - elem(car.direction, 0)}

      %Car{
        position: new_position,
        direction: {-elem(car.direction, 1), -elem(car.direction, 0)},
        next_turn: car.next_turn
      }
    end

    def move(car, "\\") do
      new_position =
        {elem(car.position, 0) + elem(car.direction, 1),
         elem(car.position, 1) + elem(car.direction, 0)}

      %Car{
        position: new_position,
        direction: {elem(car.direction, 1), elem(car.direction, 0)},
        next_turn: car.next_turn
      }
    end

    def move(car, "+") do
      case car.next_turn do
        :left -> move(car, "+", :left)
        :ahead -> move(car, "+", :ahead)
        :right -> move(car, "+", :right)
      end
    end

    def move(car, "+", :left) do
      turned_car = %Car{
        position: car.position,
        direction: {elem(car.direction, 1), -elem(car.direction, 0)},
        next_turn: :ahead
      }

      move(turned_car, "|")
    end

    def move(car, "+", :ahead) do
      turned_car = %{car | next_turn: :right}
      move(turned_car, "|")
    end

    def move(car, "+", :right) do
      turned_car = %Car{
        position: car.position,
        direction: {-elem(car.direction, 1), elem(car.direction, 0)},
        next_turn: :left
      }

      move(turned_car, "|")
    end
  end

  def first_crash(input) do
    {cars, track} = parse_track(input)
    first_crash(cars, track, 0)
  end

  def last_car_standing(input) do
    {cars, track} = parse_track(input)
    last_car_standing(cars, track, 0)
  end

  def first_crash(cars, track, iteration) do
    sorted_cars = Enum.sort_by(cars, & &1.position)

    {status, new_cars} =
      Enum.reduce_while(sorted_cars, {:start, []}, fn car_before_move, {_, acc} ->
        new_car = Car.move(car_before_move, track[car_before_move.position])

        cond do
          Enum.any?(acc, fn car -> car.position == car_before_move.position end) ->
            {:halt, {{:crash, car_before_move.position}, []}}

          Enum.any?(acc, fn car -> car.position == new_car.position end) ->
            {:halt, {{:crash, new_car.position}, []}}

          true ->
            {:cont, {:no_crash, [new_car | acc]}}
        end
      end)

    if status == :no_crash do
      first_crash(new_cars, track, iteration + 1)
    else
      elem(status, 1)
    end
  end

  def last_car_standing(cars, track, iteration) do
    sorted_cars = Enum.sort_by(cars, & &1.position)

    new_cars =
      Enum.reduce(sorted_cars, [], fn car_before_move, cars_after_move ->
        car_after_move = Car.move(car_before_move, track[car_before_move.position])

        if crashing_before_or_after_move(car_before_move, car_after_move, cars_after_move) do
          Enum.reject(cars_after_move, fn car ->
            car.position == car_before_move.position || car.position == car_after_move.position
          end)
        else
          [car_after_move | cars_after_move]
        end
      end)

    if length(new_cars) == 1 do
      Enum.at(new_cars, 0).position
    else
      last_car_standing(new_cars, track, iteration + 1)
    end
  end

  defp crashing_before_or_after_move(car_before, car_after, other_cars) do
    Enum.any?(other_cars, fn car ->
      car.position == car_before.position || car.position == car_after.position
    end)
  end

  defp parse_track(track) do
    track
    |> Enum.with_index()
    |> Enum.flat_map(fn {row, row_number} ->
      String.codepoints(row)
      |> Enum.with_index()
      |> Enum.map(fn {symbol, col_number} ->
        {row_number, col_number, symbol}
      end)
    end)
    |> Enum.reduce({[], %{}}, fn {y, x, symbol}, {cars, track} ->
      parse_symbol(symbol, {x, y}, cars, track)
    end)
  end

  defp parse_symbol(symbol, {x, y}, cars, track) when symbol in ["-", "|", "/", "\\", "+"] do
    {cars, Map.put(track, {x, y}, symbol)}
  end

  defp parse_symbol("^", {x, y}, cars, track) do
    {[%Car{position: {x, y}, direction: {0, -1}} | cars], Map.put(track, {x, y}, "|")}
  end

  defp parse_symbol("v", {x, y}, cars, track) do
    {[%Car{position: {x, y}, direction: {0, 1}} | cars], Map.put(track, {x, y}, "|")}
  end

  defp parse_symbol("<", {x, y}, cars, track) do
    {[%Car{position: {x, y}, direction: {-1, 0}} | cars], Map.put(track, {x, y}, "-")}
  end

  defp parse_symbol(">", {x, y}, cars, track) do
    {[%Car{position: {x, y}, direction: {1, 0}} | cars], Map.put(track, {x, y}, "-")}
  end

  defp parse_symbol(_, _, cars, track) do
    {cars, track}
  end
end
