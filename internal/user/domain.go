package user

import "time"

type User struct {
	Id        string     `json:"id" gorm:"type:char(36);not_null;primary_key;unique_index"`
	FirstName string     `json:"first_name" gorm:"type:varchar(100);not_null"`
	LastName  string     `json:"last_name" gorm:"type:varchar(100);not_null"`
	Email     string     `json:"email" gorm:"type:varchar(100);not_null"`
	Phone     string     `json:"phone" gorm:"type:varchar(30);not_null"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}
