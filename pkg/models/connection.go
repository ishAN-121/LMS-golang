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
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", config.Config.DB_USERNAME, config.Config.DB_PASSWORD, config.Config.DB_HOST, config.Config.DB_NAME)
}
func Connection() (*sql.DB, error) {  
	db, err := sql.Open("mysql", dsn()) 
		if err != nil {
		
		log.Printf("Error: %s when opening DB", err)
		return nil, err
		}
		db.SetMaxOpenConns(20)
		db.SetMaxIdleConns(20)
		db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancelfunc()
    err = db.PingContext(ctx)
    if err != nil {
        log.Printf("Errors %s pinging DB", err)
        return nil, err
    }
	return db, err
}