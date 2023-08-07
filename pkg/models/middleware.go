 package models

import(
	"log"
)

func Middleware(cookieid string) (string ,bool , bool) {
	db, err := Connection() 
	if err != nil {
		log.Printf("Error: %s when opening DB", err)
		}
	
	var userId int 
	var sessionid string 
	var username string
	var admin bool

	query := "SELECT userId, sessionId, username FROM cookies WHERE sessionid = ?;"
	err = db.QueryRow(query, cookieid).Scan(&userId,&sessionid,&username)
	if err != nil {
		return "",false,false
	}else{
		query :="SELECT admin FROM users WHERE username = ?"
		err := db.QueryRow(query, username).Scan(&admin)
		if err != nil {
			log.Println(err)
		}
		return username,true,admin
	}

}