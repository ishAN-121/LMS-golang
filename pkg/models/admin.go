package models
import(
	"log"
	"LMS/pkg/types"
)

func AddNewbook(title,author string , copies int) (types.Error,error){
	var error types.Error
	db, err := Connection()

	if err != nil {
	log.Printf("Error connecting to database")
	return error,err
	}
	query := "SELECT EXISTS (SELECT 1 FROM books WHERE title = ? AND author = ?)"
	var exists bool
	err = db.QueryRow(query,title,author).Scan(&exists)
	if err != nil {
		log.Println(err)
		return error,err
	}
	if exists{
		error.Msg = "Book Already exists"
		return error,err
	}else{
		query := "INSERT INTO books (author,title,count,totalcount) VALUES (?, ?, ?, ?)"
		_ = db.QueryRow(query,author,title,copies,copies)
		error.Msg = "Added Successfully"
		return error,err
	}
}

func Addbook(title,author string , copies int) (types.Error,error){
	var error types.Error
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
		return error,err
	}
	
		query := "UPDATE books SET count = count + ?, totalcount = totalcount + ? WHERE author = ? AND title= ? "
		_ = db.QueryRow(query,copies,copies,author,title)
		error.Msg = "Added Successfully"
		return error,err
	
}

func Deletebook(title,author string , copies int) (types.Error,error){
	var error types.Error
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
		return error,err
	}
	var count int
	query := "SELECT count FROM books WHERE title = ? AND author = ?"
	err = db.QueryRow(query,title,author).Scan(&count)
	if err != nil {
		log.Println(err)
		return error,err
	}
	
	if (count != 0 ){
		if(copies > count){
			error.Msg = "Too many copies for deletion"
			return error,err
		}else{
			query := "UPDATE books SET count = count - ?, totalcount = totalcount - ? WHERE author = ? AND title = ? "
			
			_ = db.QueryRow(query,copies,copies,author,title)
			error.Msg = "Deletion Successful"
			return error,err
		}

	}else{
		error.Msg = "Book not available for deletion"
		return error,err
	}
}


func Requestedbooks() (types.RequestLists,error){
	db, err := Connection()
	var requestlist types.RequestLists

	if err != nil {
		log.Printf("Error connecting to database")
		return requestlist,err
	}
	query := "SELECT id,username,bookId FROM requests WHERE status = 'requested'"
	var requests []types.Request

	rows , err := db.Query(query)
	if err != nil {
		log.Println(err)
		return requestlist,err
	}
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.Id , &request.Username,&request.Bookid)
		if err != nil {
			log.Println(err)
			return requestlist,err
		}
		requests = append(requests, request)
	}
	defer rows.Close()
	
	requestlist.Requests = requests
	return requestlist,err
}


func Approvecheckout(Id string) (types.Error,error) {
	db, err := Connection()
	var msg types.Error
	if err != nil {
		log.Printf("Error connecting to database")
		return msg,err
	}
	var bookId,count int 
	query := "SELECT bookId from requests WHERE id = ?"
	err = db.QueryRow(query,Id).Scan(&bookId)
	if err != nil {
		log.Println(err)
		return msg,err
	}else{
		query = "SELECT count FROM books WHERE id = ?"
		err = db.QueryRow(query,bookId).Scan(&count)
		if err != nil {
			log.Println(err)
			return msg,err
	}else{
		if count == 0 {
			msg.Msg = "Book not available to checkout"
			return msg,err
		}else{
			query = "UPDATE requests SET status = 'owned' WHERE id = ?"
			_, err = db.Exec(query,Id)
			if err != nil {
			log.Println(err)
			return msg,err
			}else{
			query = "UPDATE books SET count = count - 1 WHERE id = ?"
			_,err = db.Exec(query,bookId)
			if err != nil {
				log.Println(err)
				return msg,err
				}
			}
		}
	}
}
return msg,err
}
	


func Denycheckout(Id string) (types.Error,error) {
	db, err := Connection()
	var msg types.Error
	msg.Msg = ""
	if err != nil {
		log.Printf("Error connecting to database")
		return msg,err
	}
	query := "UPDATE requests SET status = NULL WHERE id = ?"
	_, err = db.Exec(query,Id)
	if err != nil{
		log.Println(err)
		return msg,err
	}
	return msg,err
}



func Checkedinbooks() (types.RequestLists,error){
	db, err := Connection()
	var requestlist types.RequestLists

	if err != nil {
		log.Printf("Error connecting to database")
		return requestlist,err
	}
	query := "SELECT id,username,bookId FROM requests WHERE status = 'checkin'"
	var requests []types.Request

	rows , err := db.Query(query)
	if err != nil {
		log.Println(err)
		return requestlist,err
	}
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.Id , &request.Username,&request.Bookid)
		if err != nil {
			log.Println(err)
			return requestlist,err
		}
		requests = append(requests, request)
	}
	defer rows.Close()
	
	requestlist.Requests = requests
	return requestlist,err
}


func Approvecheckin(Id string)error{
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
		return err
	}
	query := "UPDATE requests SET status = 'returned' WHERE id = ?"
	_,err = db.Exec(query,Id)
	if err != nil {
		log.Println(err)
		return err
	}else{
		query := "UPDATE books SET count = count + 1 WHERE id IN (SELECT bookId FROM requests WHERE id= ?)"
		_,err = db.Exec(query,Id)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return err
}


func Denycheckin(Id string)error{
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
		return err
	}

	query := "UPDATE requests SET status = 'owned' WHERE id = ?"
	_,err = db.Exec(query,Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func AdminRequestUserIds() (types.Userlist,error){
	db, err := Connection()
	var userList types.Userlist
	if err != nil {
		log.Printf("Error connecting to database")
		return userList,err
	}
	query := "SELECT id,username FROM users WHERE adminrequest = 1"
	var users []types.User

	rows , err := db.Query(query)
	if err != nil {
		log.Println(err)
		return userList,err
	}
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.Id , &user.Username)
		if err != nil {
			log.Println(err)
			return userList,err
		}
		users = append(users, user)
	}
	defer rows.Close()
	
	userList.Users = users
	return userList,err
}


func Approveadminrequest(Id string)error{
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
		return err
	}
	query := "UPDATE users SET admin = 1, adminrequest = 0 WHERE id = ?"
	_,err = db.Exec(query,Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}

func Denyadminrequest(Id string)error{
	db, err := Connection()
	log.Println(Id)
	if err != nil {
		log.Printf("Error connecting to database")
		return err
	}
	query := "UPDATE users SET adminrequest = 0 WHERE id = ?"
	_,err = db.Exec(query,Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}