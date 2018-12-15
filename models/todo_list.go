package models

// TodoList an todolist
type TodoList struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Todos       []Todo `json:"todos"`
}
