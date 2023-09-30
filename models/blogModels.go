package models

type Blog struct {
	BlogID  int    `gorm:"primary_key;auto_increment;" json:"blog_id"`
	UserID  int    `json:"user_id"`
	Title   string `gorm:"size:255;null;" json:"title"`
	Type    string `gorm:"size:255;null;" json:"type"`
	Content string `gorm:"type:LONGTEXT;null;" json:"content"`
	// Picture int    `json:"picture"`
	Status  bool `gorm:"default:true;" json:"status"`
	IsEvent bool `gorm:"default:false;" json:"is_event"`
	BaseTime
}

type BlogSelect struct {
	BlogID  int    `gorm:"primary_key;auto_increment;" json:"blog_id"`
	UserID  int    `json:"user_id"`
	Title   string `gorm:"size:255;null;" json:"title"`
	Type    string `gorm:"size:255;null;" json:"type"`
	Content string `gorm:"type:LONGTEXT;null;" json:"content"`
	// Picture int    `json:"picture"`
	Status  bool `gorm:"default:true;" json:"status"`
	IsEvent bool `gorm:"default:false;" json:"is_event"`

	User User `gorm:"foreignkey:UserID"`
	BaseTime
}
