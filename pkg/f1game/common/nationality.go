package common

type Nationality struct {
	ID   int
	Name string
}

func NewNationality(id int, name string) Nationality {
	return Nationality{id, name}
}
