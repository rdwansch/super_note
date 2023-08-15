package repository

import (
	"context"
	"database/sql"
	"fmt"
	"rdwansch/super_note/domain"
)

type noteRepository struct {
	*sql.DB
}

func NewNoteRepository(db *sql.DB) domain.NoteRepository {
	return &noteRepository{db}
}

func (n *noteRepository) FindAll(ctx context.Context) (notes []domain.Note) {
	rows, _ := n.DB.QueryContext(ctx, "SELECT id, title, content, cover FROM notes")

	for rows.Next() {
		note := domain.Note{}
		rows.Scan(&note.Id, &note.Title, &note.Content, &note.Cover)
		notes = append(notes, note)
	}

	return notes
}

func (n *noteRepository) FindById(ctx context.Context, id int) (domain.Note, error) {
	var note domain.Note

	err := n.DB.QueryRowContext(ctx, "SELECT id, title, content, cover FROM notes WHERE id = ?", id).
		Scan(&note.Id, &note.Title, &note.Content, &note.Cover)

	return note, err

}

func (n *noteRepository) Create(ctx context.Context, note *domain.Note) (err error) {
	stmt, err := n.DB.PrepareContext(ctx, "INSERT INTO notes (title, content, cover, id_user) VALUES (?,?,?,?)")

	if err != nil {
		return
	}
	res, err := stmt.ExecContext(ctx, note.Title, note.Content, note.Cover, note.Author.Id)

	if err != nil {
		return
	}
	id, err := res.LastInsertId()
	note.Id = int(id)

	return
}

func (n *noteRepository) Delete(ctx context.Context, id int) (err error) {
	stmt, err := n.DB.PrepareContext(ctx, "DELETE FROM notes WHERE id = ?")

	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, id)
	if err != nil {
		return
	}

	rowAffected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowAffected < 1 {
		err = fmt.Errorf("something wrong with code, total affected: %d", rowAffected)
		return
	}

	return

}
