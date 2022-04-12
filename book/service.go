package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(Id int) (Book, error)
	Create(book BookRequest) (Book, error)
}
type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}
func (s *service) FindAll() ([]Book, error) {
	//books, err := s.repository.FindAll()
	//return books,err
	return s.repository.FindAll()
}
func (s *service) FindByID(Id int) (Book, error) {
	book, err := s.repository.FindByID(Id)
	return book, err
}

// func (s *service) Create(request BookRequest BookRequest)(Book, error){
// 	book := Book{
// 		Title: bookRe
// 	}
// }
