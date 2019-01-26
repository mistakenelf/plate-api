defmodule PlateAPIWeb.Schema.DataTypes do
  use Absinthe.Schema.Notation

  object :task do
    field :name, :string
    field :description, :string
  end
end
