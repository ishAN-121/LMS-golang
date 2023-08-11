package models

import (
	"log"

	"LMS/pkg/types"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)


func GetBooks(db *sql.DB) (types.ListBooks,error) {

	query := "SELECT * FROM books"
	rows, err := db.Query(query)
	db.Close()

	var listBooks types.ListBooks
	if err != nil {
		log.Printf("error %s querying the database", err)
		return listBooks,err
	}


	var fetchBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Id, &book.Title,&book.Author,&book.Copies,&book.Totalcount)
		if err != nil {
			log.Printf("error %s scanning the row", err)
			return listBooks,err
		}
		fetchBooks = append(fetchBooks, book)
	}

	
	listBooks.Books = fetchBooks
	return listBooks,err

}