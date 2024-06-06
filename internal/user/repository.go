package user

import (
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(user *User) error
		GetAll() ([]User, error)
		Get(id string) (*User, error)
		// Update(id string, user *User) error
		// Delete(id string) error
	}

	repository struct {
		log *log.Logger
		db  *gorm.DB
	}
)

func NewRepository(log *log.Logger, db *gorm.DB) Repository {
	return &repository{
		log: log,
		db:  db,
	}
}

func (repo *repository) Create(user *User) error {
	user.Id = uuid.New().String()

	if err := repo.db.Create(user).Error; err != nil {
		repo.log.Println(err)
		return err
	}

	repo.log.Println("user created with id: ", user.Id)
	return nil
}

func (repo *repository) GetAll() ([]User, error) {
	var users []User

	// if err := repo.db.Model(&users).Select("id, first_name, email, created_at").Order("created_at desc").Find(&users).Error; err != nil {
	if err := repo.db.Model(&users).Order("created_at desc").Find(&users).Error; err != nil {
		repo.log.Println(err)
		return nil, err
	}

	return users, nil
}

func (repo *repository) Get(id string) (*User, error) {
	user := User{
		Id: id,
	}

	// if err := repo.db.Model(&user).Where("id = ?", id).First(&user).Error; err != nil {
	if err := repo.db.Model(&user).First(&user).Error; err != nil {
		repo.log.Println(err)
		return nil, err
	}

	return &user, nil
}
