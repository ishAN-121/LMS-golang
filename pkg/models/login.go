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
	db, err := Connection()

	var msg types.Error

	if err != nil {
		msg.Msg = "Error in connecting to database"
		return false, msg, err
	}
	defer db.Close()

	query := "SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)"

	var exists bool
	err = db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		log.Println(err)
		return false, msg, err
	}
	if exists {
		var hashedpass string 
		var admin bool

		query := "SELECT hash,admin FROM users WHERE username = ?"

		err = db.QueryRow(query, username).Scan(&hashedpass , &admin)
		if err != nil {
		log.Println(err)
		return false, msg, err
		}
		err = bcrypt.CompareHashAndPassword([]byte(hashedpass), []byte(password))
		if err != nil {
		msg.Msg = "wrong password"
		return false, msg, err
		}else{
			session_id := uuid.New().String()
			cookie := http.Cookie{
				Name:     "SessionID",
				Value:    session_id,
				Expires:  time.Now().Add(24 * time.Hour), 
				HttpOnly: true,                           
			}
			http.SetCookie(w, &cookie)
			var id int
			query := "SELECT id FROM users WHERE username = ?"
			err = db.QueryRow(query, username).Scan(&id)
			if err != nil {
				log.Println(err)
				return false, msg, err
			}
			db.Exec("INSERT INTO cookies (sessionId, userId, username) VALUES (?, ?,?)", session_id, id,username)
			
			msg.Msg = "Login successful"
			return admin,msg,err
		}
	}else{
		msg.Msg = "Username does not exist"
		return false,msg , err
	}
}

