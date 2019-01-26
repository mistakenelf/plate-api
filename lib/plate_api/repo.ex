defmodule PlateAPI.Repo do
  use Ecto.Repo,
    otp_app: :plate_api,
    adapter: Ecto.Adapters.Postgres
end
