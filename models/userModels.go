package models

type User struct {
	UserID   int    `gorm:"primary_key" json:"user_id"`
	BlogID   string `json:"blog_id"`
	Username string `json:"username"`
	Email    string `gorm:"unique" json:"email"`
	Password string `json:"password"`
	Role     string `gorm:"type:ENUM('super_admin', 'admin', 'user');not null" json:"role"`
	BaseTime
}
