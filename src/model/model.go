// src/model/model.go
package model

import (
    "database/sql"
    "log"
    // "os"

    // mysql driver
    _ "github.com/go-sql-driver/mysql"
)

// DBConnect returns *sql.DB
func DBConnect() (db *sql.DB) {
    dbDriver := "mysql"
    dbUser := "go_user"
    dbPass := "wataru"
    dbName := "go_web_app"
    dbOption := "?parseTime=true"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName+dbOption)
    if err != nil {
        log.Fatal(err)
    }
    return db
}
