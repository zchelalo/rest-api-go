package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Id        string         `json:"id" gorm:"type:char(36);not null;primary_key"`
	FirstName string         `json:"first_name" gorm:"type:varchar(100);not null"`
	LastName  string         `json:"last_name" gorm:"type:varchar(100);not null"`
	Email     string         `json:"email" gorm:"type:varchar(100);not null;unique"`
	Phone     string         `json:"phone" gorm:"type:varchar(30);not null"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
	if user.Id == "" {
		user.Id = uuid.New().String()
	}
	return
}
