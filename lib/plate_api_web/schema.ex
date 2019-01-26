defmodule PlateAPIWeb.Schema do
  use Absinthe.Schema

  import_types(PlateAPIWeb.Schema.DataTypes)

  query do
    @desc "Get list of tasks"
    field :tasks, list_of(:task) do
      resolve(fn _parent, _args, _resolution ->
        {:ok, PlateAPI.Task.list_tasks()}
      end)
    end
  end
end
