package models

import(
	"LMS/pkg/types"
	"log"
	"time"
	"net/http"
	

	"golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

func Authenticate(w http.ResponseWriter, r *http.Request,username , password string) (bool , types.Error,error){
	db, error := Connection()

	var message types.Error

	if error != nil {
		message.Message = "Error in connecting to database"
		return false, message, error
	}
	defer db.Close()

	query := "SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)"

	var exists bool
	error = db.QueryRow(query, username).Scan(&exists)
	if error != nil {
		log.Println(error)
		return false, message, error
	}
	if exists {
		var hashedPass string 
		var admin bool

		query := "SELECT hash,admin FROM users WHERE username = ?"

		error = db.QueryRow(query, username).Scan(&hashedPass , &admin)
		if error != nil {
		log.Println(error)
		return false, message, error
		}
		error = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))
		if error != nil {
		message.Message = "wrong password"
		return false, message, error
		}else{
			sessionId := uuid.New().String()
			cookie := http.Cookie{
				Name:     "sessionId",
				Value:    sessionId,
				Expires:  time.Now().Add(24 * time.Hour), 
				HttpOnly: true,                           
			}
			http.SetCookie(w, &cookie)
			var id int
			query := "SELECT id FROM users WHERE username = ?"
			error = db.QueryRow(query, username).Scan(&id)
			if error != nil {
				log.Println(error)
				return false, message, error
			}
			db.Exec("INSERT INTO cookies (sessionId, userId, username) VALUES (?, ?,?)", sessionId, id,username)
			
			message.Message = "Login successful"
			return admin,message,error
		}
	}else{
		message.Message = "Username does not exist"
		return false,message , error
	}
}

