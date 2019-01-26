defmodule PlateAPIWeb.Router do
  use PlateAPIWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api" do
    pipe_through :api
    forward "/graphiql", Absinthe.Plug.GraphiQL, schema: PlateAPIWeb.Schema
  end
end
