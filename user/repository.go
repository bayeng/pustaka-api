package user

import "gorm.io/gorm"

type Repository interface {
	FindUsers() ([]User, error)
	FindUser(id int) (User, error)
	CreateUser(user User) (User, error)
	DeleteUser(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]User, error) {
	var users []User

	err := r.db.Find(&users).Error

	return users, err

}

func (r *repository) FindUser(id int) (User, error) {
	var user User
	err := r.db.First(&user, id).Error

	return user, err
}

func (r *repository) CreateUser(user User) (User, error) {

	err := r.db.Create(&user).Error
	return user, err
}

func (r *repository) DeleteUser(user User) (User, error) {

	err := r.db.Delete(&user).Error

	return user, err
}
