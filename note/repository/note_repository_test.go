package repository

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestNoteRepositoryFindAll(t *testing.T) {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/super_note")

	if err != nil {
		panic("Error on connection SQL: " + err.Error())
	}

	err = db.Ping()

	if err != nil {
		panic("Error on PING connection: " + err.Error())
	}

	noteRepository := NewNoteRepository(db)

	notes := noteRepository.FindAll()

	fmt.Println(notes)
}
