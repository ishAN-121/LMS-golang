package models

import (
	"LMS/pkg/types"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, error := bcrypt.GenerateFromPassword([]byte(password), 10)
	if error != nil {
		return "", error
	}
	return string(hash), nil
}

func AddUser(username, password string) (types.Error,error) {
	var message types.Error
	db, error := Connection()

	if error != nil {
		log.Printf("Error connecting to database")
		return message,error
	}


	query := "SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)"

	var exists bool
	error = db.QueryRow(query, username).Scan(&exists)
	if error != nil {
		log.Println(error)
		return message,error
	}

	if exists {
		message.Msg = "User already exists"
		return message,error
	} else {
		hashedPassword, error := hashPassword(password)
		if error != nil {
			log.Println(error)
			return message,error
		}
		_, error = db.Exec(`INSERT INTO users (username,hash,admin,adminrequest) VALUES (?, ?, ?,?)`, username, hashedPassword, 0, 0)
		if error != nil {
			log.Println(error)
		} else {
			message.Msg = "Registered Successfully"
			return message,error
		}
		return message,error
	}
}
