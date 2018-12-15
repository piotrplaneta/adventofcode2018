defmodule Day15 do
  defmodule Elf do
    defstruct position: {-1, -1}, type: :elf, health: 300
  end

  defmodule Gremlin do
    defstruct position: {-1, -1}, type: :gremlin, health: 300
  end

  defmodule Wall do
    defstruct position: {-1, -1}, type: :wall
  end

  defmodule EmptySpace do
    defstruct position: {-1, -1}, type: :empty_space
  end

  defmodule PathToObject do
    defstruct object: nil, length: -1, steps: []
  end

  def movement_tick({elfs, gremlins, game_state}) do
    players = elfs ++ gremlins
    sorted_players = Enum.sort_by(players, &{elem(&1.position, 1), elem(&1.position, 0)})

    {new_players, _, new_game_state} =
      sorted_players
      |> Enum.reduce({[], sorted_players, game_state}, fn player,
                                                          {moved_players, old_players,
                                                           current_game_state} ->
        {moved_player, new_game_state} =
          if player.type == :elf do
            move(player, gremlins(moved_players ++ old_players), current_game_state)
          else
            move(player, elfs(moved_players ++ old_players), current_game_state)
          end

        {moved_players ++ [moved_player], Enum.drop(old_players, 1), new_game_state}
      end)

    {elfs(new_players), gremlins(new_players), new_game_state}
  end

  defp move(player, enemies, game_state) do
    closest_enemy_path = calculate_closest_enemy_path(player, enemies, game_state)

    if closest_enemy_path == [] do
      {player, game_state}
    else
      selected_move =
        Enum.min_by([{0, -1}, {-1, 0}, {1, 0}, {0, 1}], fn move ->
          new_coordinates =
            {elem(player.position, 0) + elem(move, 0), elem(player.position, 1) + elem(move, 1)}

          if game_state[new_coordinates].type != :empty_space do
            :infinity
          else
            shortest_path_length(
              %Elf{position: new_coordinates},
              %EmptySpace{
                position: Enum.at(closest_enemy_path, -1)
              },
              game_state
            )
          end
        end)

      new_coordinates =
        {elem(player.position, 0) + elem(selected_move, 0),
         elem(player.position, 1) + elem(selected_move, 1)}

      new_player =
        if player.type == :elf do
          %Elf{position: new_coordinates, health: player.health}
        else
          %Gremlin{position: new_coordinates, health: player.health}
        end

      new_game_state =
        game_state
        |> Map.put(player.position, %EmptySpace{position: player.position})
        |> Map.put(new_coordinates, new_player)

      {new_player, new_game_state}
    end
  end

  defp calculate_closest_enemy_path(player, enemies, game_state) do
    destination_points =
      Enum.flat_map(enemies, fn enemy -> empty_adjacent_points(enemy, game_state) end)

    paths = Enum.map(destination_points, fn point -> shortest_path(player, point, game_state) end)

    Enum.filter(paths, fn path -> path != [] end)
    |> Enum.map(fn path -> Enum.drop(path, 1) end)
    |> Enum.sort_by(&{length(&1), elem(Enum.at(&1, 0), 1), elem(Enum.at(&1, 0), 0)})
    |> Enum.at(0)
  end

  defp shortest_path_length(player, destination, game_state) do
    path = shortest_path(player, destination, game_state)

    if path == [] do
      :infinity
    else
      length(path)
    end
  end

  defp shortest_path(player, destination, game_state) do
    start_point = player.position
    end_point = destination.position

    bfs([[start_point]], end_point, MapSet.new(), game_state)
  end

  defp bfs([], _, _, _), do: []

  defp bfs(queue, end_point, visited, game_state) do
    path = Enum.at(queue, 0)

    current_point = Enum.at(path, -1)

    if current_point == end_point do
      path
    else
      adjacent_points =
        Enum.map(
          empty_adjacent_points(%EmptySpace{position: current_point}, game_state),
          fn x -> x.position end
        )

      {new_queue, new_visited} =
        Enum.reduce(adjacent_points, {queue, visited}, fn adjacent_point,
                                                          {queue_so_far, visited_so_far} ->
          if !MapSet.member?(visited, adjacent_point) do
            new_path = path ++ [adjacent_point]
            {queue_so_far ++ [new_path], MapSet.put(visited_so_far, adjacent_point)}
          else
            {queue_so_far, visited_so_far}
          end
        end)

      bfs(Enum.drop(new_queue, 1), end_point, new_visited, game_state)
    end
  end

  def empty_adjacent_points(object, game_state) do
    [{0, -1}, {-1, 0}, {1, 0}, {0, 1}]
    |> Enum.map(fn difference ->
      game_state[
        {elem(object.position, 0) + elem(difference, 0),
         elem(object.position, 1) + elem(difference, 1)}
      ]
    end)
    |> Enum.filter(fn point ->
      case point.type do
        :empty_space -> true
        _ -> false
      end
    end)
  end

  def parse_input(input) do
    game_state =
      Enum.with_index(input)
      |> Enum.reduce(%{}, fn {row, row_index}, objects ->
        String.codepoints(row)
        |> Enum.with_index()
        |> Enum.reduce(objects, fn {object, column_index}, row_objects ->
          Map.put(
            row_objects,
            {column_index, row_index},
            parse_object(object, column_index, row_index)
          )
        end)
      end)

    {elfs(Map.values(game_state)), gremlins(Map.values(game_state)), game_state}
  end

  defp parse_object("#", x, y), do: %Wall{position: {x, y}}
  defp parse_object(".", x, y), do: %EmptySpace{position: {x, y}}
  defp parse_object("E", x, y), do: %Elf{position: {x, y}}
  defp parse_object("G", x, y), do: %Gremlin{position: {x, y}}

  defp elfs(list) do
    Enum.filter(list, fn object ->
      object.type == :elf
    end)
  end

  defp gremlins(list) do
    Enum.filter(list, fn object ->
      object.type == :gremlin
    end)
  end
end
