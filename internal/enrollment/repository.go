package enrollment

import (
	"log"

	"github.com/zchelalo/rest-api-go/internal/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(enrollment *domain.Enrollment) error
	}

	repository struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepository(log *log.Logger, db *gorm.DB) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}

func (repo *repository) Create(enrollment *domain.Enrollment) error {
	if err := repo.db.Create(enrollment).Error; err != nil {
		repo.log.Printf("error: %v", err)
		return err
	}

	repo.log.Println("enrollment created with id: ", enrollment.Id)
	return nil
}
