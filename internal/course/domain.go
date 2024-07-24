package course

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	Id        string         `json:"id" gorm:"type:char(36);not null;primary_key"`
	Name      string         `json:"name" gorm:"type:varchar(50);not null"`
	StartDate time.Time      `json:"start_date" gorm:"not null"`
	EndDate   time.Time      `json:"end_date" gorm:"not null"`
	CreatedAt *time.Time     `json:"-"`
	UpdatedAt *time.Time     `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
}

func (course *Course) BeforeCreate(tx *gorm.DB) (err error) {
	if course.Id == "" {
		course.Id = uuid.New().String()
	}
	return
}
