package models

// TodoList an todolist
type TodoList struct {
	ID          string `json:"id" gorm:"default:uuid_generate_v4();primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Todos       []Todo `json:"todos" gorm:"many2one:todos"`
}
