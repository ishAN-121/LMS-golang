package models

import (
	"LMS/pkg/types"
	"fmt"
	"log"
	"strings"

)

func Checkout(username string, bookId int) (types.Error , error) {
	var message types.Error
	db, error := Connection()

	if error != nil {
		log.Printf("Error connecting to database")
		return message,error
	}
		var exists bool
		query := "SELECT 1 FROM requests WHERE bookId = ? AND username = ? AND (status = 'requested' OR status = 'owned')"
		error = db.QueryRow(query, bookId, username).Scan(&exists)
		fmt.Println(error)
		if exists {
			message.Msg = "Already Requested or Owned"
			return message,error
		} else {
			query := "INSERT INTO requests (bookId , username , status) VALUES (?,?, 'requested');"
			_, error = db.Exec(query, bookId, username)
			if error != nil {
				log.Println(error)
			}  else {
					message.Msg = "Book checked out"
					return message,error
				}
			}
			return message,error
		}
		

func Checkin(username string, id int) error {
	db, error := Connection()

	if error != nil {
		log.Printf("Error connecting to database")
		return error
	}
	query := "UPDATE requests SET status = 'checkin' WHERE  bookId= ? AND username = ? AND status = 'owned' "
	_, error = db.Exec(query, id, username)
	if error != nil {
		log.Println(error)
		return error
	}	
	return error
} 
	
func IssuedBooks(username string) (types.ListBooks,error) {
	
	var listBooks types.ListBooks
	db, error := Connection()
	if error != nil {
		log.Printf("error %s connecting to the database", error)
		return listBooks,error
	}
	query := "SELECT bookId FROM requests WHERE username = ? AND status = 'owned'"
	rows, error := db.Query(query,username)

	if error != nil {
		log.Printf("error %s querying the database", error)
		return listBooks,error
	}
	defer rows.Close()
	var fetchBooks []types.Book
	
	for rows.Next() {
		var book types.Book
		error := rows.Scan(&book.Id)
		if error != nil {
			log.Printf("error %s scanning the row", error)
			return listBooks,error
		}
		fetchBooks = append(fetchBooks, book)
	}
	

	if len(fetchBooks) == 0 {
		listBooks.Books = fetchBooks
		return listBooks,error
	}

	bookIDsStr := make([]string, len(fetchBooks))
	for i, book:= range fetchBooks {
		bookIDsStr[i] = book.Id
	}

	rows, error = db.Query(fmt.Sprintf("SELECT * FROM books WHERE id IN (%s)", strings.Join(bookIDsStr, ",")))
	if error != nil {
		log.Println(error)
		return listBooks,error
	}
    var fetchIssuedBooks []types.Book
	for rows.Next() {
		var book types.Book
		error := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Copies, &book.Totalcount)
		if error != nil {
			log.Println(error)
			return listBooks,error
		}
		fetchIssuedBooks = append(fetchIssuedBooks, book)
	}
	listBooks.Books = fetchIssuedBooks
	return listBooks,error
}

func AdminRequest(username string)error{
	db, error := Connection()
	if error != nil {
		log.Printf("error %s connecting to the database", error)
		return error
	}
	query := "UPDATE users SET adminrequest = 1 WHERE username = ? "
	_,error = db.Exec(query, username)
	if error != nil {
		log.Println(error)
		return error
	}
	return error
}