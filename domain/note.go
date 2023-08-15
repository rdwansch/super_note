package domain

type Note struct {
	Id      int    `json:"id"`
	Tite    string `json:"title"`
	Content string `json:"content"`
	Cover   string `json:"cover"`

	Author *Author // belongs to
}

type NoteRepository interface {
	FindAll() (notes []Note)
	FindById(id int) (Note, error)
	Create(note Note) error
	Delete(id int) error
}
