package common

type Team struct {
	ID   int
	Name string
}

func NewTeam(id int, name string) Team {
	return Team{id, name}
}
