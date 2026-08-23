package common

type Track struct {
	ID   int
	Name string
}

func NewTrack(id int, name string) Track {
	return Track{id, name}
}
