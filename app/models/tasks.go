package models

// Task describes a task
type Task struct {
	ID        string `json:"id" gorm:"default:uuid_generate_v4();primary_key"`
	Name      string `json:"name"`
	DueDate   string `json:"dueDate"`
	Template  string `json:"template"`
	Status    string `json:"status"`
	Content   string `json:"content"`
	CreatedBy string `json:"createdBy"`
}
