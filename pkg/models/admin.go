package models
import(
	"log"
	"LMS/pkg/types"
)

func Addbook(title,author string , copies int) types.Error{
	var error types.Error
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
	}
	query := "SELECT EXISTS (SELECT 1 FROM books WHERE title = ? AND author = ?)"
	var exists bool
	err = db.QueryRow(query,title,author).Scan(&exists)
	if err != nil {
		log.Println(err)
	}
	if exists{
		query := "UPDATE books SET count = count + ?, totalcount = totalcount + ? WHERE author = ? AND title= ? "
		_ = db.QueryRow(query,copies,copies,author,title)
		error.Msg = "Added Successfully"
		return error
	}else{
		query := "INSERT INTO books (author,title,count,totalcount) VALUES (?, ?, ?, ?)"
		_ = db.QueryRow(query,author,title,copies,copies)
		error.Msg = "Added Successfully"
		return error
	}
}

func Deletebook(title,author string , copies int) types.Error{
	var error types.Error
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
	}
	var count int
	query := "SELECT count FROM books WHERE title = ? AND author = ?"
	err = db.QueryRow(query,title,author).Scan(&count)
	if err != nil {
		log.Println(err)
		error.Msg = "Book does not exist"
		return error
	}
	
	if (count != 0 ){
		if(copies > count){
			error.Msg = "Too many copies for deletion"
			return error
		}else{
			query := "UPDATE books SET count = count - ?, totalcount = totalcount - ? WHERE author = ? AND title = ? "
			
			_ = db.QueryRow(query,copies,copies,author,title)
			error.Msg = "Deletion Successful"
			return error
		}

	}else{
		error.Msg = "Book not available for deletion"
		return error
	}
}