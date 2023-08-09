package models

import (
	"LMS/pkg/types"
	"fmt"
	"log"
	"strings"
)

func Checkout(username string, bookid int) types.Error {
	var error types.Error
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
	}
		var exists bool
		query := "SELECT 1 FROM requests WHERE bookId = ? AND username = ? AND (status = 'requested' OR status = 'owned')"
		err = db.QueryRow(query, bookid, username).Scan(&exists)
		fmt.Println(err)
		if exists {
			error.Msg = "Already Requested or Owned"
			
			return error
		} else {
			query := "INSERT INTO requests (bookId , username , status) VALUES (?,?, 'requested');"
			_, err = db.Exec(query, bookid, username)
			if err != nil {
				log.Println(err)
			}  else {
					error.Msg = "Book checked out"
					return error
				}
			}
			return error
		}
		

func Checkin(username string, id int) {
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
	}
	query := "UPDATE requests SET status = 'checkin' WHERE  bookId= ? AND username = ? AND status = 'owned' "
	_, err = db.Exec(query, id, username)
	if err != nil {
		log.Println(err)
	}	
} 
	
func Issuedbooks(username string) types.ListBooks {

	db, err := Connection()
	if err != nil {
		log.Printf("error %s connecting to the database", err)
	}
	query := "SELECT bookId FROM requests WHERE username = ? AND status = 'owned'"
	rows, err := db.Query(query,username)

	if err != nil {
		log.Printf("error %s querying the database", err)
	}
	defer rows.Close()
	var fetchBooks []types.Book
	var listBooks types.ListBooks
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Id)
		if err != nil {
			log.Printf("error %s scanning the row", err)
		}
		fetchBooks = append(fetchBooks, book)
	}
	

	if len(fetchBooks) == 0 {
		listBooks.Books = fetchBooks
		return listBooks
	}

	bookIDsStr := make([]string, len(fetchBooks))
	for i, book:= range fetchBooks {
		bookIDsStr[i] = book.Id
	}

	

	rows, err = db.Query(fmt.Sprintf("SELECT * FROM books WHERE id IN (%s)", strings.Join(bookIDsStr, ",")))
	if err != nil {
		log.Println(err)
	}
    var fetchissuedBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Copies, &book.Totalcount)
		if err != nil {
			log.Println(err)
		}
		fetchissuedBooks = append(fetchissuedBooks, book)
	}
	listBooks.Books = fetchissuedBooks
	return listBooks
}

func Adminrequest(username string){
	db, err := Connection()
	if err != nil {
		log.Printf("error %s connecting to the database", err)
	}
	query := "UPDATE users SET adminrequest = 1 WHERE username = ? "
	_,err = db.Exec(query, username)
	if err != nil {
		log.Println(err)
	}
}