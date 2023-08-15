package repository

import (
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
}
