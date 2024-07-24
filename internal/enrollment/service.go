package enrollment

import (
	"errors"
	"log"

	"github.com/zchelalo/rest-api-go/internal/course"
	"github.com/zchelalo/rest-api-go/internal/domain"
	"github.com/zchelalo/rest-api-go/internal/user"
)

type (
	Service interface {
		Create(userId, courseId string) (*domain.Enrollment, error)
	}

	service struct {
		log           *log.Logger
		userService   user.Service
		courseService course.Service
		repository    Repository
	}
)

func NewService(repo Repository, log *log.Logger, userService user.Service, courseService course.Service) Service {
	return &service{
		repository:    repo,
		log:           log,
		userService:   userService,
		courseService: courseService,
	}
}

func (srv service) Create(userId, courseId string) (*domain.Enrollment, error) {
	enrollment := &domain.Enrollment{
		UserId:   userId,
		CourseId: courseId,
		Status:   "P",
	}

	if _, err := srv.userService.Get(enrollment.UserId); err != nil {
		return nil, errors.New("user id doesn't exists")
	}

	if _, err := srv.courseService.Get(enrollment.CourseId); err != nil {
		return nil, errors.New("course id doesn't exists")
	}

	if err := srv.repository.Create(enrollment); err != nil {
		srv.log.Println(err)
		return nil, err
	}

	return enrollment, nil
}
