package models

// User current user object
type User struct {
	ID           string `json:"id" gorm:"default:uuid_generate_v4();primary_key"`
	Email        string `json:"email" gorm:"type:varchar(100);unique"`
	PasswordHash string `json:"passwordHash"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}

// LoginRequest login input
type LoginRequest struct {
	Email    string
	Password string
}

// RegisterRequest regisger input
type RegisterRequest struct {
	Email     string
	Password  string
	FirstName string
	LastName  string
}
