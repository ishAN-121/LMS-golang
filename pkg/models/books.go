package models

import (
	"log"

	"LMS/pkg/types"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)


func GetBooks(db *sql.DB) (types.ListBooks,error) {

	query := "SELECT * FROM books"
	rows, error := db.Query(query)
	db.Close()

	var listBooks types.ListBooks
	if error != nil {
		log.Printf("error %s querying the database", error)
		return listBooks,error
	}


	var fetchBooks []types.Book
	for rows.Next() {
		var book types.Book
		error := rows.Scan(&book.Id, &book.Title,&book.Author,&book.Copies,&book.Totalcount)
		if error != nil {
			log.Printf("error %s scanning the row", error)
			return listBooks,error
		}
		fetchBooks = append(fetchBooks, book)
	}

	
	listBooks.Books = fetchBooks
	return listBooks,error

}