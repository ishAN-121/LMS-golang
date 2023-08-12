package models

import (
	"LMS/pkg/types"
	"fmt"
	"log"
	"strings"

)

func Checkout(username string, bookId int) (types.Error , error) {
	var msg types.Error
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
		return msg,err
	}
		var exists bool
		query := "SELECT 1 FROM requests WHERE bookId = ? AND username = ? AND (status = 'requested' OR status = 'owned')"
		err = db.QueryRow(query, bookId, username).Scan(&exists)
		fmt.Println(err)
		if exists {
			msg.Msg = "Already Requested or Owned"
			return msg,err
		} else {
			query := "INSERT INTO requests (bookId , username , status) VALUES (?,?, 'requested');"
			_, err = db.Exec(query, bookId, username)
			if err != nil {
				log.Println(err)
			}  else {
					msg.Msg = "Book checked out"
					return msg,err
				}
			}
			return msg,err
		}
		

func Checkin(username string, id int) error {
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
		return err
	}
	query := "UPDATE requests SET status = 'checkin' WHERE  bookId= ? AND username = ? AND status = 'owned' "
	_, err = db.Exec(query, id, username)
	if err != nil {
		log.Println(err)
		return err
	}	
	return err
} 
	
func IssuedBooks(username string) (types.ListBooks,error) {
	
	var listBooks types.ListBooks
	db, err := Connection()
	if err != nil {
		log.Printf("error %s connecting to the database", err)
		return listBooks,err
	}
	query := "SELECT bookId FROM requests WHERE username = ? AND status = 'owned'"
	rows, err := db.Query(query,username)

	if err != nil {
		log.Printf("error %s querying the database", err)
		return listBooks,err
	}
	defer rows.Close()
	var fetchBooks []types.Book
	
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Id)
		if err != nil {
			log.Printf("error %s scanning the row", err)
			return listBooks,err
		}
		fetchBooks = append(fetchBooks, book)
	}
	

	if len(fetchBooks) == 0 {
		listBooks.Books = fetchBooks
		return listBooks,err
	}

	bookIDsStr := make([]string, len(fetchBooks))
	for i, book:= range fetchBooks {
		bookIDsStr[i] = book.Id
	}

	rows, err = db.Query(fmt.Sprintf("SELECT * FROM books WHERE id IN (%s)", strings.Join(bookIDsStr, ",")))
	if err != nil {
		log.Println(err)
		return listBooks,err
	}
    var fetchIssuedBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Copies, &book.Totalcount)
		if err != nil {
			log.Println(err)
			return listBooks,err
		}
		fetchIssuedBooks = append(fetchIssuedBooks, book)
	}
	listBooks.Books = fetchIssuedBooks
	return listBooks,err
}

func AdminRequest(username string)error{
	db, err := Connection()
	if err != nil {
		log.Printf("error %s connecting to the database", err)
		return err
	}
	query := "UPDATE users SET adminrequest = 1 WHERE username = ? "
	_,err = db.Exec(query, username)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}