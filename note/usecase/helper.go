package usecase

import "rdwansch/super_note/domain"

func ToNoteResponse(note domain.Note) domain.NoteResponse {
	return domain.NoteResponse{
		Id:      note.Id,
		Title:   note.Title,
		Content: note.Content,
		Cover:   note.Cover,
		Author: &domain.AuthorResponse{
			Name:     note.Author.Name,
			Username: note.Author.Username,
		},
	}
}

func FromNoteRequest(note *domain.NoteRequest) *domain.Note {
	return &domain.Note{
		Id:      0,
		Title:   note.Title,
		Content: note.Content,
		Cover:   note.Cover,
		Author: &domain.Author{
			Username: note.AuthorUsername,
		},
	}
}
