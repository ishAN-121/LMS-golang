package models

import(
"log"

)

func Logout(username string){
	db,err := Connection()
	if err != nil {
		log.Printf("Error: %s when opening DB", err)
		}
	query := `DELETE FROM cookies WHERE username = ?;`
	_ ,err = db.Exec(query, username)
	if err != nil {
		log.Println(db.Ping().Error())
	}
	
}