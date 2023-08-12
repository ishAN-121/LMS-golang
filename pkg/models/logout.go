package models

import(
"log"

)

func Logout(username string)error{
	db,error := Connection()
	if error != nil {
		log.Printf("Error: %s when opening DB", error)
		return error
		}
	query := `DELETE FROM cookies WHERE username = ?;`
	_ ,error = db.Exec(query, username)
	if error != nil {
		log.Println(error)
		return error
	}
	return error
}