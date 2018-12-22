package models

// Todo an todo item
type Todo struct {
	ID          string `json:"id" gorm:"default:uuid_generate_v4();primary_key"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}
