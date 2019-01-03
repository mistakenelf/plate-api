package models

// TodoList an todolist
type TodoList struct {
	ID        string `json:"id" gorm:"default:uuid_generate_v4();primary_key"`
	Title     string `json:"title"`
	Todos     []Todo `json:"todos" gorm:"foreignkey:TodoListID"`
	CreatedBy string `json:"createdBy"`
	Completed bool   `json:"completed"`
}

// Todo an todo item
type Todo struct {
	ID          string `json:"id" gorm:"default:uuid_generate_v4();primary_key"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	TodoListID  string `json:"todoListId"`
	CreatedBy   string `json:"createdBy"`
}
