package repository

import (
	"rdwansch/super_note/domain"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	_ "github.com/go-sql-driver/mysql"
)

func TestNoteRepositoryFindAll(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "content", "cover"}).
		AddRow(1, "My Birthday", "I want a Macbook Pro", nil).
		AddRow(2, "Her", "When I first saw you // yeah that's a song by Beyoncé Ft. Jamie Foxx", nil)

	mock.ExpectQuery("SELECT id, title, content, cover FROM notes").WillReturnRows(rows)

	noteRepository := NewNoteRepository(db)
	notes := noteRepository.FindAll()

	assert.NotEmpty(t, notes)
	assert.NotNil(t, notes)
	assert.Len(t, notes, 2)
}

func TestNoteRepositoryFindById(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "content", "cover"}).
		AddRow(1, "My Birthday", "I want a Macbook Pro", "")

	mock.ExpectQuery("SELECT id, title, content, cover FROM notes WHERE id = ?").
		WithArgs(1).WillReturnRows(rows)

	noteRepository := NewNoteRepository(db)
	notes, err := noteRepository.FindById(1)

	assert.NoError(t, err)
	assert.Nil(t, err)
	assert.Equal(t, "My Birthday", notes.Title)
}

func TestNoteRepositoryCreate(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	notes := &domain.Note{
		Title:   "My Birthday",
		Content: "I want a Macbook Pro",
		Cover:   "",
		Author: &domain.Author{
			Id: 1,
		},
	}

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectPrepare("INSERT INTO notes (title, content, cover, id_user) VALUES (?,?,?,?)").
		ExpectExec().WithArgs(notes.Title, notes.Content, notes.Cover, notes.Author.Id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	noteRepository := NewNoteRepository(db)
	err = noteRepository.Create(notes)

	assert.NoError(t, err)
	assert.Equal(t, 1, notes.Id)
}

func TestRepositoryDelete(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectPrepare("DELETE FROM notes WHERE id = ?").
		ExpectExec().WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	noteRepository := NewNoteRepository(db)
	err = noteRepository.Delete(1)

	assert.NoError(t, err)
	assert.Nil(t, err)
}
