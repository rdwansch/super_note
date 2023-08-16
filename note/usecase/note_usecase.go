package usecase

import (
	"context"
	"rdwansch/super_note/domain"
)

type NoteUsecase struct {
	domain.NoteRepository
}

func (n *NoteUsecase) FindAll(ctx context.Context) (note []domain.NoteResponse) {
	notes := n.NoteRepository.FindAll(ctx)

	var noteResponses []domain.NoteResponse
	for _, note := range notes {
		noteResponses = append(noteResponses, ToNoteResponse(note))
	}

	return noteResponses
}

func (n *NoteUsecase) FindById(ctx context.Context, id int) (note domain.NoteResponse, err error) {
	res, err := n.NoteRepository.FindById(ctx, id)

	if err != nil {
		return
	}

	return ToNoteResponse(res), err
}

func (n *NoteUsecase) Create(ctx context.Context, note *domain.NoteRequest) (err error) {
	return n.NoteRepository.Create(ctx, FromNoteRequest(note))
}

func (n *NoteUsecase) Delete(ctx context.Context, id int) (err error) {
	return n.NoteRepository.Delete(ctx, id)
}
