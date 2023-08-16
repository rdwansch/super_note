package domain

import "context"

// Represent database structure
type Note struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Cover   string `json:"cover"`

	Author *Author // belongs to
}

// Represent response API
type NoteResponse struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Cover   string `json:"cover"`

	Author *AuthorResponse
}

type NoteRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Cover   string `json:"cover"`

	AuthorUsername string
}

type NoteRepository interface {
	FindAll(ctx context.Context) (notes []Note)
	FindById(ctx context.Context, id int) (Note, error)
	Create(ctx context.Context, note *Note) (err error)
	Delete(ctx context.Context, id int) (err error)
}

type NoteUsecase interface {
	FindAll(ctx context.Context) (note []NoteResponse)
	FindById(ctx context.Context, id int) (note NoteResponse, err error)
	Create(ctx context.Context, note *NoteRequest) (err error)
	Delete(ctx context.Context, id int) (err error)
}
