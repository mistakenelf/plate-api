defmodule PlateAPI.Task.Tasks do
  use Ecto.Schema
  import Ecto.Changeset

  schema "tasks" do
    field :description, :string
    field :name, :string

    timestamps()
  end

  @doc false
  def changeset(tasks, attrs) do
    tasks
    |> cast(attrs, [:name, :description])
    |> validate_required([:name, :description])
  end
end
