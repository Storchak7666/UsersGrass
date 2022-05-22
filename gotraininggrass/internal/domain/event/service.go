package event

type Service interface {
	FindAll() ([]User, error)
	FindByName(name string) ([]User, error)
	CreateUser(name string, age int64, city string, country string) (*User, error)
	UpdateById(id int64, name string, age int64, city string, country string) (*User, error)
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]User, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindByName(name string) ([]User, error) {
	return (*s.repo).FindByName(name)
}

func (s *service) CreateUser(name string, age int64, city string, country string) (*User, error) {
	return (*s.repo).CreateUser(name, age, city, country)
}

func (s *service) UpdateById(id int64, name string, age int64, city string, country string) (*User, error) {
	return (*s.repo).UpdateById(id, name, age, city, country)
}
