	package models
	import(
		"log"
		"LMS/pkg/types"
	)

	func AddNewbook(title,author string , copies int) types.Error{
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
			error.Msg = "Book Already exists"
			return error
		}else{
			query := "INSERT INTO books (author,title,count,totalcount) VALUES (?, ?, ?, ?)"
			_ = db.QueryRow(query,author,title,copies,copies)
			error.Msg = "Added Successfully"
			return error
		}
	}

	func Addbook(title,author string , copies int) types.Error{
		var error types.Error
		db, err := Connection()

		if err != nil {
			log.Printf("Error connecting to database")
		}
		
			query := "UPDATE books SET count = count + ?, totalcount = totalcount + ? WHERE author = ? AND title= ? "
			_ = db.QueryRow(query,copies,copies,author,title)
			error.Msg = "Added Successfully"
			return error
		
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


	func Requestedbooks() types.RequestLists{
		db, err := Connection()

		if err != nil {
			log.Printf("Error connecting to database")
		}
		query := "SELECT id,username,bookId FROM requests WHERE status = 'requested'"
		var requests []types.Request

		rows , err := db.Query(query)
		if err != nil {
			log.Println(err)
		}
		for rows.Next() {
			var request types.Request
			err := rows.Scan(&request.Id , &request.Username,&request.Bookid)
			if err != nil {
				log.Println(err)
			}
			requests = append(requests, request)
		}
		defer rows.Close()
		var requestlist types.RequestLists
		requestlist.Requests = requests
		return requestlist
	}


	func Approvecheckout(Id string) types.Error {
		db, err := Connection()
		if err != nil {
			log.Printf("Error connecting to database")
		}
		var bookId,count int 
		var error types.Error
		query := "SELECT bookId from requests WHERE id = ?"
		err = db.QueryRow(query,Id).Scan(&bookId)
		if err != nil {
			log.Println(err)
		}else{
			query = "SELECT count FROM books WHERE id = ?"
			err = db.QueryRow(query,bookId).Scan(&count)
			if err != nil {
				log.Println(err)
		}else{
			if count == 0 {
				error.Msg = "Book not available to checkout"
				return error
			}else{
				query = "UPDATE requests SET status = 'owned' WHERE id = ?"
				_, err = db.Exec(query,Id)
				if err != nil {
				log.Println(err)
				}else{
				query = "UPDATE books SET count = count - 1 WHERE id = ?"
				_,err = db.Exec(query,bookId)
				if err != nil {
					log.Println(err)
					}
				}
			}
		}
	}
	return error
}
		


	func Denycheckout(Id string) types.Error {
		db, err := Connection()
		var error types.Error
		error.Msg = ""
		if err != nil {
			log.Printf("Error connecting to database")
		}
		query := "UPDATE requests SET status = NULL WHERE id = ?"
		_, err = db.Exec(query,Id)
		if err != nil{
			log.Println(err)
		}
		return error
	}



	func Checkedinbooks() types.RequestLists{
		db, err := Connection()

		if err != nil {
			log.Printf("Error connecting to database")
		}
		query := "SELECT id,username,bookId FROM requests WHERE status = 'checkin'"
		var requests []types.Request

		rows , err := db.Query(query)
		if err != nil {
			log.Println(err)
		}
		for rows.Next() {
			var request types.Request
			err := rows.Scan(&request.Id , &request.Username,&request.Bookid)
			if err != nil {
				log.Println(err)
			}
			requests = append(requests, request)
		}
		defer rows.Close()
		var requestlist types.RequestLists
		requestlist.Requests = requests
		log.Println(requestlist)
		return requestlist
	}


	func Approvecheckin(Id string){
		db, err := Connection()

		if err != nil {
			log.Printf("Error connecting to database")
		}
		query := "UPDATE requests SET status = 'returned' WHERE id = ?"
		_,err = db.Exec(query,Id)
		if err != nil {
			log.Println(err)
		}else{
			query := "UPDATE books SET count = count + 1 WHERE id IN (SELECT bookId FROM requests WHERE id= ?)"
			_,err = db.Exec(query,Id)
			if err != nil {
				log.Println(err)
			}
		}
	}


	func Denycheckin(Id string){
		db, err := Connection()

		if err != nil {
			log.Printf("Error connecting to database")
		}
		query := "UPDATE requests SET status = 'owned' WHERE id = ?"
		_,err = db.Exec(query,Id)
		if err != nil {
			log.Println(err)
		}
	}

	func AdminRequestUserIds() types.Userlist{
		db, err := Connection()

		if err != nil {
			log.Printf("Error connecting to database")
		}
		query := "SELECT id,username FROM users WHERE adminrequest = 1"
		var users []types.User

		rows , err := db.Query(query)
		if err != nil {
			log.Println(err)
		}
		for rows.Next() {
			var user types.User
			err := rows.Scan(&user.Id , &user.Username)
			if err != nil {
				log.Println(err)
			}
			users = append(users, user)
		}
		defer rows.Close()
		var userList types.Userlist
		userList.Users = users
		return userList
	}


	func Approveadminrequest(Id string){
		db, err := Connection()

		if err != nil {
			log.Printf("Error connecting to database")
		}
		query := "UPDATE users SET admin = 1, adminrequest = 0 WHERE id = ?"
		_,err = db.Exec(query,Id)
		if err != nil {
			log.Println(err)
		}
	}

	func Denyadminrequest(Id string){
		db, err := Connection()
		log.Println(Id)
		if err != nil {
			log.Printf("Error connecting to database")
		}
		query := "UPDATE users SET adminrequest = 0 WHERE id = ?"
		_,err = db.Exec(query,Id)
		if err != nil {
			log.Println(err)
		}
	}