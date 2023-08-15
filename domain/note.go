package domain

import "context"

type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Cover   string `json:"cover"`

	Author *Author // belongs to
}

type NoteRepository interface {
	FindAll(ctx context.Context) (notes []Note)
	FindById(ctx context.Context, id int) (Note, error)
	Create(ctx context.Context, note *Note) (err error)
	Delete(ctx context.Context, id int) (err error)
}
