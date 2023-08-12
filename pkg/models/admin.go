package models
import(
	"log"
	"LMS/pkg/types"
)

func AddNewBook(title,author string , copies int) (types.Error,error){
	var message types.Error
	db, error := Connection()

	if error != nil {
	log.Printf("Error connecting to database")
	return message,error
	}
	query := "SELECT EXISTS (SELECT 1 FROM books WHERE title = ? AND author = ?)"
	var exists bool
	error = db.QueryRow(query,title,author).Scan(&exists)
	if error != nil {
		log.Println(error)
		return message,error
	}
	if exists{
		message.Message = "Book Already exists"
		return message,error
	}else{
		query := "INSERT INTO books (author,title,count,totalcount) VALUES (?, ?, ?, ?)"
		_ = db.QueryRow(query,author,title,copies,copies)
		message.Message = "Added Successfully"
		return message,error
	}
}

func AddBook(title,author string , copies int) (types.Error,error){
	var message types.Error
	db, error := Connection()

	if error != nil {
		log.Printf("Error connecting to database")
		return message,error
	}
	
	query := "UPDATE books SET count = count + ?, totalcount = totalcount + ? WHERE author = ? AND title= ? "
	_ = db.QueryRow(query,copies,copies,author,title)
	message.Message = "Added Successfully"
	return message,error
}

func DeleteBook(title,author string , copies int) (types.Error,error){
	var message types.Error
	db, error := Connection()

	if error != nil {
		log.Printf("Error connecting to database")
		return message,error
	}
	var count int
	query := "SELECT count FROM books WHERE title = ? AND author = ?"
	error = db.QueryRow(query,title,author).Scan(&count)
	if error != nil {
		log.Println(error)
		return message,error
	}
	
	if (count != 0 ){
		if(copies > count){
			message.Message = "Too many copies for deletion"
			return message,error
		}else{
			query := "UPDATE books SET count = count - ?, totalcount = totalcount - ? WHERE author = ? AND title = ? "
			
			_ = db.QueryRow(query,copies,copies,author,title)
			message.Message = "Deletion Successful"
			return message,error
		}

	}else{
		message.Message = "Book not available for deletion"
		return message,error
	}
}


func RequestedBooks() (types.RequestLists,error){
	db, error := Connection()
	var requestList types.RequestLists

	if error != nil {
		log.Printf("Error connecting to database")
		return requestList,error
	}
	query := "SELECT id,username,bookId FROM requests WHERE status = 'requested'"
	var requests []types.Request

	rows , error := db.Query(query)
	if error != nil {
		log.Println(error)
		return requestList,error
	}
	for rows.Next() {
		var request types.Request
		error := rows.Scan(&request.Id , &request.Username,&request.BookId)
		if error != nil {
			log.Println(error)
			return requestList,error
		}
		requests = append(requests, request)
	}
	defer rows.Close()
	
	requestList.Requests = requests
	return requestList,error
}


func ApproveCheckout(Id string) (types.Error,error) {
	db, error := Connection()
	var message types.Error
	if error != nil {
		log.Printf("Error connecting to database")
		return message,error
	}
	var bookId,count int 
	query := "SELECT bookId from requests WHERE id = ?"
	error = db.QueryRow(query,Id).Scan(&bookId)
	if error != nil {
		log.Println(error)
		return message,error
	}else{
		query = "SELECT count FROM books WHERE id = ?"
		error = db.QueryRow(query,bookId).Scan(&count)
		if error != nil {
			log.Println(error)
			return message,error
	}else{
		if count == 0 {
			message.Message = "Book not available to checkout"
			return message,error
		}else{
			query = "UPDATE requests SET status = 'owned' WHERE id = ?"
			_, error = db.Exec(query,Id)
			if error != nil {
			log.Println(error)
			return message,error
			}else{
			query = "UPDATE books SET count = count - 1 WHERE id = ?"
			_,error = db.Exec(query,bookId)
			if error != nil {
				log.Println(error)
				return message,error
				}
			}
		}
	}
}
return message,error
}
	


func DenyCheckout(Id string) (types.Error,error) {
	db, error := Connection()
	var message types.Error
	message.Message = ""
	if error != nil {
		log.Printf("Error connecting to database")
		return message,error
	}
	query := "UPDATE requests SET status = NULL WHERE id = ?"
	_, error = db.Exec(query,Id)
	if error != nil{
		log.Println(error)
		return message,error
	}
	return message,error
}



func CheckedinBooks() (types.RequestLists,error){
	db, error := Connection()
	var requestList types.RequestLists

	if error != nil {
		log.Printf("Error connecting to database")
		return requestList,error
	}
	query := "SELECT id,username,bookId FROM requests WHERE status = 'checkin'"
	var requests []types.Request

	rows , error := db.Query(query)
	if error != nil {
		log.Println(error)
		return requestList,error
	}
	for rows.Next() {
		var request types.Request
		error := rows.Scan(&request.Id , &request.Username,&request.BookId)
		if error != nil {
			log.Println(error)
			return requestList,error
		}
		requests = append(requests, request)
	}
	defer rows.Close()
	
	requestList.Requests = requests
	return requestList,error
}


func ApproveCheckin(Id string)error{
	db, error := Connection()

	if error != nil {
		log.Printf("Error connecting to database")
		return error
	}
	query := "UPDATE requests SET status = 'returned' WHERE id = ?"
	_,error = db.Exec(query,Id)
	if error != nil {
		log.Println(error)
		return error
	}else{
		query := "UPDATE books SET count = count + 1 WHERE id IN (SELECT bookId FROM requests WHERE id= ?)"
		_,error = db.Exec(query,Id)
		if error != nil {
			log.Println(error)
			return error
		}
	}
	return error
}


func DenyCheckin(Id string)error{
	db, error := Connection()

	if error != nil {
		log.Printf("Error connecting to database")
		return error
	}

	query := "UPDATE requests SET status = 'owned' WHERE id = ?"
	_,error = db.Exec(query,Id)
	if error != nil {
		log.Println(error)
		return error
	}
	return error
}

func AdminRequestUserIds() (types.Userlist,error){
	db, error := Connection()
	var userList types.Userlist
	if error != nil {
		log.Printf("Error connecting to database")
		return userList,error
	}
	query := "SELECT id,username FROM users WHERE adminrequest = 1"
	var users []types.User

	rows , error := db.Query(query)
	if error != nil {
		log.Println(error)
		return userList,error
	}
	for rows.Next() {
		var user types.User
		error := rows.Scan(&user.Id , &user.Username)
		if error != nil {
			log.Println(error)
			return userList,error
		}
		users = append(users, user)
	}
	defer rows.Close()
	
	userList.Users = users
	return userList,error
}


func ApproveAdminRequest(Id string)error{
	db, error := Connection()

	if error != nil {
		log.Printf("Error connecting to database")
		return error
	}
	query := "UPDATE users SET admin = 1, adminrequest = 0 WHERE id = ?"
	_,error = db.Exec(query,Id)
	if error != nil {
		log.Println(error)
		return error
	}
	return error
}

func DenyAdminRequest(Id string)error{
	db, error := Connection()
	log.Println(Id)
	if error != nil {
		log.Printf("Error connecting to database")
		return error
	}
	query := "UPDATE users SET adminrequest = 0 WHERE id = ?"
	_,error = db.Exec(query,Id)
	if error != nil {
		log.Println(error)
		return error
	}
	return error
}