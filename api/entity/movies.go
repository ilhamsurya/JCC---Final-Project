package entity

import "time"

// Movie struct
type Movie struct {
	ID          int       `json:"id" gorm:"primary_key;auto_increment"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Duration    int       `json:"duration"`
	Genre       string    `json:"genre"`
	Image_url   string    `json:"image_url"`
	Rating      int       `json:"rating"`
	Review      string    `json:"review"`
	Director    *Director `json:"director"`
	Cast        *Cast     `json:"cast"`
	Year        int       `json:"year"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:createdAt"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"default:null"`
}
