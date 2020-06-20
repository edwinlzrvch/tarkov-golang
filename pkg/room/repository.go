package room

type Reader interface {
	Find(id string) (*Room, error)
}

type Writer interface {
	Add(room *Room) (*Room, error)
	// Delete(id string) error
}

type Repository interface {
	Reader
	Writer
}
