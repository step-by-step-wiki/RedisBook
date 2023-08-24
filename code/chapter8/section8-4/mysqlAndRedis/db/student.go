package db

type Student struct {
	Id    int
	Name  string
	Age   int
	Score float64
}

func (s *Student) FindById(id int) error {
	return Conn.Where("id = ?", id).First(s).Error
}
