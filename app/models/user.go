package models

// User current user object
type User struct {
	ID           string `json:"id" gorm:"default:uuid_generate_v4()"`
	Email        string `json:"email" gorm:"type:varchar(100);unique"`
	PasswordHash string `json:"passwordHash"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
}
