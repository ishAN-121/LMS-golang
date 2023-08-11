package models

import(
"log"

)

func Logout(username string)error{
	db,err := Connection()
	if err != nil {
		log.Printf("Error: %s when opening DB", err)
		return err
		}
	query := `DELETE FROM cookies WHERE username = ?;`
	_ ,err = db.Exec(query, username)
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}