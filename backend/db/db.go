package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	lDBUser := "STN013"
	lDBPassword := "Best@123"
	lDBName := "training"
	lDBHost := "192.168.2.5"

	lDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", lDBUser, lDBPassword, lDBHost, lDBName)

	lDB, lErr := sql.Open("mysql", lDSN)
	if lErr != nil {
		log.Println("DBConnect refused:", lErr)
		return nil, lErr
	}

	lDB.SetMaxOpenConns(20)
	lDB.SetMaxIdleConns(10)
	lDB.SetConnMaxIdleTime(60 * time.Second)

	lErr = lDB.Ping()
	if lErr != nil {
		log.Println("error while pinging database:", lErr)
		return nil, lErr
	}

	return lDB, nil
}
