package user

import "time"

type User struct {
	Id        string     `json:"id" gorm:"type:char(36);not null;primary_key"`
	FirstName string     `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName  string     `json:"last_name" gorm:"type:varchar(100);not null"`
	Email     string     `json:"email" gorm:"type:varchar(100);not null;unique"`
	Phone     string     `json:"phone" gorm:"type:varchar(30);not null"`
	CreatedAt *time.Time `json:"-"`
	UpdatedAt *time.Time `json:"-"`
}
