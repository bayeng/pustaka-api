package user

type Service interface {
	FindUsers() ([]User, error)
	FindUser(id int) (User, error)
	CreateUser(userInput InputUser) (User, error)
	DeleteUser(id int) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindUsers() ([]User, error) {

	book, err := s.repository.FindUsers()
	return book, err
}

func (s *service) FindUser(id int) (User, error) {

	user, err := s.repository.FindUser(id)
	return user, err
}

func (s *service) CreateUser(userInput InputUser) (User, error) {

	user := User{
		Name:     userInput.Name,
		Username: userInput.Username,
		Password: userInput.Password,
	}

	newUser, err := s.repository.CreateUser(user)

	return newUser, err
}

func (s *service) DeleteUser(id int) (User, error) {

	user, _ := s.repository.FindUser(id)

	delUser, err := s.repository.DeleteUser(user)

	return delUser, err
}
