defmodule Day8 do
  @moduledoc """
  Solutions for day8
  """

  @doc """
  Sum of tree metadata for a tree described with a string.

  ## Examples

      iex> Day8.sum_tree_metadata("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2")
      138

  """
  def sum_tree_metadata(input) do
    input |> parse_tree() |> traverse_tree_for_metadata_sum() |> elem(0)
  end

  @doc """
  Tree root node value

  ## Examples

      iex> Day8.calculate_root_node_value("2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2")
      66

  """
  def calculate_root_node_value(input) do
    input |> parse_tree() |> traverse_tree_for_node_value() |> elem(0)
  end

  defp traverse_tree_for_metadata_sum([child_count, metadata_count | rest]) do
    {children_metadata_sum, rest_of_description} =
      Enum.reduce(replicate_trues(child_count), {0, rest}, fn _, {meta_sum, rest_of_desc} ->
        {child_metadata_sum, after_child_description} = traverse_tree_for_metadata_sum(rest_of_desc)
        {meta_sum + child_metadata_sum, after_child_description}
      end)

    node_metadata_sum = rest_of_description |> Enum.take(metadata_count) |> Enum.sum()
    rest_of_description = Enum.drop(rest_of_description, metadata_count)

    {node_metadata_sum + children_metadata_sum, rest_of_description}
  end

  defp traverse_tree_for_node_value([child_count, metadata_count | rest]) do
    {child_values, rest_of_description} =
      Enum.reduce(replicate_trues(child_count), {[], rest}, fn _, {child_vals, rest_of_desc} ->
        {child_value, after_child_description} = traverse_tree_for_node_value(rest_of_desc)
        {child_vals ++ [child_value], after_child_description}
      end)

    node_value = cond do
      length(child_values) == 0 ->
        rest_of_description |> Enum.take(metadata_count) |> Enum.sum()
      length(child_values) > 0 ->
        rest_of_description
        |> Enum.take(metadata_count)
        |> Enum.map(&(Enum.at(child_values, &1 - 1, 0)))
        |> Enum.sum()
    end

    rest_of_description = Enum.drop(rest_of_description, metadata_count)
    {node_value, rest_of_description}
  end

  defp parse_tree(string) do
    string |> String.split(" ") |> Enum.map(&String.to_integer/1)
  end

  def replicate_trues(n), do: for i <- 0..n, i > 0, do: true
end
