package book

type Service interface {
	FindAllBook() ([]Book, error)
	FindById(id int) (Book, error)
	CreateBook(book BookInput) (Book, error)
	UpdateBook(id int, bookUpdate BookInput) (Book, error)
	DeleteBook(id int) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAllBook() ([]Book, error) {

	books, err := s.repository.FindAllBook()
	return books, err
}

func (s *service) FindById(id int) (Book, error) {

	book, err := s.repository.FindById(id)
	return book, err
}

func (s *service) CreateBook(bookInput BookInput) (Book, error) {

	price, _ := bookInput.Price.Int64()
	rating, _ := bookInput.Rating.Int64()
	book := Book{
		Title:  bookInput.Title,
		Desc:   bookInput.Desc,
		Price:  int(price),
		Rating: int(rating),
	}

	new, err := s.repository.CreateBook(book)

	return new, err
}
func (s *service) UpdateBook(id int, bookUpdate BookInput) (Book, error) {

	book, _ := s.repository.FindById(id)
	price, _ := bookUpdate.Price.Int64()
	rating, _ := bookUpdate.Rating.Int64()

	book.Title = bookUpdate.Title
	book.Desc = bookUpdate.Desc
	book.Price = int(price)
	book.Rating = int(rating)

	updateBook, err := s.repository.UpdateBook(book)

	return updateBook, err

}

func (s *service) DeleteBook(id int) (Book, error) {

	book, _ := s.repository.FindById(id)
	delBook, err := s.repository.DeleteBook(book)

	return delBook, err

}
