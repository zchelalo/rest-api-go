package course

import (
	"log"
	"time"

	"github.com/zchelalo/rest-api-go/internal/domain"
)

type (
	Filters struct {
		Name string
	}

	Service interface {
		Create(name, startDate, endDate string) (*domain.Course, error)
		GetAll(filters Filters, offset, limit int) ([]domain.Course, error)
		Get(id string) (*domain.Course, error)
		Update(id string, name, startDate, endDate *string) error
		Delete(id string) error
		Count(filters Filters) (int, error)
	}

	service struct {
		log        *log.Logger
		repository Repository
	}
)

func NewService(repo Repository, log *log.Logger) Service {
	return &service{
		repository: repo,
		log:        log,
	}
}

func (srv *service) Create(name, startDate, endDate string) (*domain.Course, error) {
	startDateParsed, err := time.Parse("2006-01-02", startDate)
	if err != nil {
		srv.log.Println(err)
		return nil, err
	}

	endDateParsed, err := time.Parse("2006-01-02", endDate)
	if err != nil {
		srv.log.Println(err)
		return nil, err
	}

	course := &domain.Course{
		Name:      name,
		StartDate: startDateParsed,
		EndDate:   endDateParsed,
	}

	if err := srv.repository.Create(course); err != nil {
		srv.log.Println(err)
		return nil, err
	}

	return course, nil
}

func (srv *service) GetAll(filters Filters, offset, limit int) ([]domain.Course, error) {
	srv.log.Println("get all courses service")
	courses, err := srv.repository.GetAll(filters, offset, limit)
	if err != nil {
		return nil, err
	}
	return courses, nil
}

func (srv *service) Get(id string) (*domain.Course, error) {
	srv.log.Println("get course service")
	course, err := srv.repository.Get(id)
	if err != nil {
		// srv.log.Println(err)
		return nil, err
	}
	return course, nil
}

func (srv *service) Update(id string, name, startDate, endDate *string) error {
	srv.log.Println("update course service")

	var startDateParsed *time.Time
	if startDate != nil {
		parsed, err := time.Parse("2006-01-02", *startDate)
		if err != nil {
			srv.log.Println(err)
			return err
		}
		startDateParsed = &parsed
	}

	var endDateParsed *time.Time
	if endDate != nil {
		parsed, err := time.Parse("2006-01-02", *endDate)
		if err != nil {
			srv.log.Println(err)
			return err
		}
		endDateParsed = &parsed
	}

	return srv.repository.Update(id, name, startDateParsed, endDateParsed)
}

func (srv *service) Delete(id string) error {
	srv.log.Println("delete course service")
	return srv.repository.Delete(id)
}

func (srv *service) Count(filters Filters) (int, error) {
	srv.log.Println("count course service")
	return srv.repository.Count(filters)
}
