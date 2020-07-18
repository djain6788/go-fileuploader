package main

import (
	"database/sql"
	"log"
	"time"
)

func dbConn() (db *sql.DB) {
	dbDriver := configuration1.dbDriver
	dbUser := configuration1.dbUser
	dbPass := configuration1.dbPass
	dbName := configuration1.dbName
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func insertFileMetaData(fmd fileMetaData) {
	db := dbConn()

	insForm, err := db.Prepare("INSERT IGNORE INTO files_metadata(file_name, mime_type, account_id,record_id,creation_time) VALUES(?,?,?,?,?)")
	if err != nil {
		panic(err.Error())
	}
	var datetime = time.Now()
	insForm.Exec(fmd.fileName, fmd.mimeType, fmd.accountId, fmd.recordId, datetime)
	log.Println("INSERT: Name: " + fmd.fileName)
	defer db.Close()

}
