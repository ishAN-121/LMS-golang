package models

import (

	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)
func TestBooks(t *testing.T) {

	db, mock, err := sqlmock.New()
	if err != nil {
    	t.Fatalf("Error creating mock database connection: %v", err)
    }
    defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "title", "author", "count","totalcount"}).
        AddRow("1", "abc", "xyz", 5, 7)

	mock.ExpectQuery("SELECT \\* FROM books").
    WillReturnRows(rows)


	booksList  := GetBooks(db)
	
	books := booksList.Books

	assert.Equal(t, 1, len(books))

	assert.Equal(t, "1", books[0].Id)
    assert.Equal(t, "abc", books[0].Title)
    assert.Equal(t, "xyz", books[0].Author)
    assert.Equal(t, 5, books[0].Copies)
	assert.Equal(t, 7, books[0].Totalcount)


	if err := mock.ExpectationsWereMet(); err != nil {
        t.Errorf("Unfulfilled expectations: %s", err)
    }
}



