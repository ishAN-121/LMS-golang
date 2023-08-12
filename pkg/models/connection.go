package models

import (
	"database/sql"

	"context"
	"fmt"
	"log"
	"time"
	

	_ "github.com/go-sql-driver/mysql"
	"LMS/config"
)

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Config.DBUSERNAME, config.Config.DBPASSWORD, config.Config.DBHOST, config.Config.DBNAME)
}
func Connection() (*sql.DB, error) {  
	db, error := sql.Open("mysql", dsn()) 
		if error != nil {
		
		log.Printf("Error: %s when opening DB", error)
		return nil, error
		}
		db.SetMaxOpenConns(20)
		db.SetMaxIdleConns(20)
		db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    error = db.PingContext(ctx)
    if error != nil {
        log.Printf("Errors %s pinging DB", error)
        return nil, error
    }
	return db, error
}