package repository

import (
	"context"
	"database/sql"
	"rdwansch/super_note/domain"
)

type noteRepository struct {
	*sql.DB
}

// db:= NewConnection()
// noteRepository:=NewNoteRepository(db)
// noteService := NewNoteService(noteRepository)
func NewNoteRepository(db *sql.DB) domain.NoteRepository {
	return &noteRepository{db}
}

func (n *noteRepository) FindAll() (notes []domain.Note) {
	defer n.DB.Close()

	ctx := context.Background()

	rows, err := n.DB.QueryContext(ctx, "SELECT id, title, content, cover FROM notes")

	if err != nil {
		panic("Error on query FindAll " + err.Error())
	}

	for rows.Next() {
		note := domain.Note{}
		rows.Scan(&note.Id, &note.Tite, &note.Content, &note.Cover)
		notes = append(notes, note)
	}

	return notes
}

func (n *noteRepository) FindById(id int) (domain.Note, error) {
	panic("not implemented") // TODO: Implement
}

func (n *noteRepository) Create(note domain.Note) error {
	panic("not implemented") // TODO: Implement
}

func (n *noteRepository) Delete(id int) error {
	panic("not implemented") // TODO: Implement
}
