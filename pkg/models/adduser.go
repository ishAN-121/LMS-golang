package models

import (
	"LMS/pkg/types"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func Adduser(username, password, confirmpassword string) types.Error {
	var error types.Error
	db, err := Connection()

	if err != nil {
		log.Printf("Error connecting to database")
	}


	query := "SELECT EXISTS (SELECT 1 FROM users WHERE username = ?)"

	var exists bool
	err = db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		log.Println(err)
	}

	if exists {
		error.Msg = "User already exists"
		return error
	} else {
		hashedpassword, err := hashPassword(password)
		if err != nil {
			log.Println(err)
		}
		_, err = db.Exec(`INSERT INTO users (username,hash,admin,adminrequest) VALUES (?, ?, ?,?)`, username, hashedpassword, 0, 0)
		if err != nil {
			log.Println(err)
		} else {
			error.Msg = "Registered Successfully"
			return error
		}
		return error
	}
}
