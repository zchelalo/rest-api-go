package user

import "log"

type (
	Service interface {
		Create(firstName, lastName, email, phone string) (*User, error)
		GetAll() ([]User, error)
		Get(id string) (*User, error)
		Update(id string, firstName, lastName, email, phone *string) error
		Delete(id string) error
	}

	service struct {
		log        *log.Logger
		repository Repository
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:        log,
		repository: repo,
	}
}

func (srv *service) Create(firstName, lastName, email, phone string) (*User, error) {
	srv.log.Println("create user service")
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
	}
	if err := srv.repository.Create(&user); err != nil {
		// srv.log.Println(err)
		return nil, err
	}
	return &user, nil
}

func (srv *service) GetAll() ([]User, error) {
	srv.log.Println("get all users service")
	users, err := srv.repository.GetAll()
	if err != nil {
		// srv.log.Println(err)
		return nil, err
	}
	return users, nil
}

func (srv *service) Get(id string) (*User, error) {
	srv.log.Println("get user service")
	user, err := srv.repository.Get(id)
	if err != nil {
		// srv.log.Println(err)
		return nil, err
	}
	return user, nil
}

func (srv *service) Update(id string, firstName, lastName, email, phone *string) error {
	srv.log.Println("update user service")
	if err := srv.repository.Update(id, firstName, lastName, email, phone); err != nil {
		// srv.log.Println(err)
		return err
	}
	return nil
}

func (srv *service) Delete(id string) error {
	srv.log.Println("delete user service")
	if err := srv.repository.Delete(id); err != nil {
		// srv.log.Println(err)
		return err
	}
	return nil
}
