package book

import "gorm.io/gorm"

type Repository interface {
	FindAllBook() ([]Book, error)
	FindById(id int) (Book, error)
	CreateBook(book Book) (Book, error)
	UpdateBook(book Book) (Book, error)
	DeleteBook(book Book) (Book, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAllBook() ([]Book, error) {

	var books []Book
	err := r.db.Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil

}

func (r *repository) FindById(id int) (Book, error) {

	var book Book
	err := r.db.First(&book, id).Error

	return book, err
}

func (r *repository) CreateBook(book Book) (Book, error) {

	err := r.db.Create(&book).Error

	return book, err
}

func (r *repository) UpdateBook(book Book) (Book, error) {

	err := r.db.Save(&book).Error

	return book, err
}

func (r *repository) DeleteBook(book Book) (Book, error) {

	err := r.db.Delete(&book).Error

	return book, err
}
