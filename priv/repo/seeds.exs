# Script for populating the database. You can run it as:
#
#     mix run priv/repo/seeds.exs
#
# Inside the script, you can read and write to any of your
# repositories directly:
#
#     PlateAPI.Repo.insert!(%PlateAPI.SomeSchema{})
#
# We recommend using the bang functions (`insert!`, `update!`
# and so on) as they will fail if something goes wrong.

alias PlateAPI.Repo
alias PlateAPI.Task.Tasks

Repo.insert!(%Tasks{
  name: "Woodworking",
  description: "Stuff I need to build"
})
