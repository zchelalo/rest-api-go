package course

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(course *Course) error
		GetAll(filters Filters, offset, limit int) ([]Course, error)
		Get(id string) (*Course, error)
		Update(id string, name *string, startDate, endDate *time.Time) error
		Delete(id string) error
		Count(filters Filters) (int, error)
	}

	repository struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepository(log *log.Logger, db *gorm.DB) Repository {
	return &repository{
		db:  db,
		log: log,
	}
}

func (repo *repository) Create(course *Course) error {
	if err := repo.db.Create(course).Error; err != nil {
		repo.log.Println(err)
		return err
	}

	repo.log.Println("course created with id: ", course.Id)
	return nil
}

func (repo *repository) GetAll(filters Filters, offset, limit int) ([]Course, error) {
	var courses []Course
	tx := repo.db.Model(&Course{})
	tx = applyFilters(tx, filters)
	tx = tx.Limit(limit).Offset(offset)
	if err := tx.Order("created_at desc").Find(&courses).Error; err != nil {
		repo.log.Println(err)
		return nil, err
	}

	return courses, nil
}

func (repo *repository) Get(id string) (*Course, error) {
	course := Course{
		Id: id,
	}

	if err := repo.db.Model(&course).First(&course).Error; err != nil {
		repo.log.Println(err)
		return nil, err
	}

	return &course, nil
}

func (repo *repository) Update(id string, name *string, startDate, endDate *time.Time) error {
	values := make(map[string]interface{})

	if name != nil {
		values["name"] = *name
	}

	if startDate != nil {
		values["start_date"] = *startDate
	}

	if endDate != nil {
		values["end_date"] = *endDate
	}

	if err := repo.db.Model(&Course{}).Where("id = ?", id).Updates(values).Error; err != nil {
		repo.log.Println(err)
		return err
	}

	return nil
}

func (repo *repository) Delete(id string) error {
	course := Course{
		Id: id,
	}

	if err := repo.db.Model(&course).Delete(&course).Error; err != nil {
		repo.log.Println(err)
		return err
	}

	return nil
}

func (repo *repository) Count(filters Filters) (int, error) {
	var count int64
	tx := repo.db.Model(&Course{})
	tx = applyFilters(tx, filters)
	if err := tx.Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func applyFilters(tx *gorm.DB, filters Filters) *gorm.DB {
	if filters.Name != "" {
		filters.Name = fmt.Sprintf("%%%s%%", strings.ToLower(filters.Name))
		tx = tx.Where("lower(name) like ?", filters.Name)
	}

	return tx
}
