 package models

import(
	"log"
)

func Middleware(cookieid string) (string ,bool , bool) {
	db, error := Connection() 
	if error != nil {
		log.Printf("Error: %s when opening DB", error)
		}
	
	var userId int 
	var sessionid string 
	var username string
	var admin bool

	query := "SELECT userId, sessionId, username FROM cookies WHERE sessionid = ?;"
	error = db.QueryRow(query, cookieid).Scan(&userId,&sessionid,&username)
	if error != nil {
		return "",false,false
	}else{
		query :="SELECT admin FROM users WHERE username = ?"
		error := db.QueryRow(query, username).Scan(&admin)
		if error != nil {
			log.Println(error)
		}
		return username,true,admin
	}

}