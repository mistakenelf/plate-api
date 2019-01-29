defmodule PlateAPIWeb.Router do
  use PlateAPIWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api" do
    pipe_through :api
    forward "/graphql", Absinthe.Plug, schema: PlateAPIWeb.Schema
    forward "/graphiql", Absinthe.Plug.GraphiQL, schema: PlateAPIWeb.Schema
  end
end
