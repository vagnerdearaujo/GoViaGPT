package main

import (
    "database/sql"
    "fmt"
    _ "github.com/denisenkom/go-mssqldb"
    _ "github.com/FirebirdSQL/firebird-go"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
)

func main() {
    var dbType string
    var connString string
    fmt.Print("Enter the type of database (sqlserver, firebird, mysql, postgres): ")
    fmt.Scan(&dbType)
    fmt.Print("Enter the connection string: ")
    fmt.Scan(&connString)
    
    db, err := sql.Open(dbType, connString)
    if err != nil {
        fmt.Println("Error opening database:", err)
        return
    }
    defer db.Close()
    
    err = db.Ping()
    if err != nil {
        fmt.Println("Error pinging database:", err)
        return
    }
    fmt.Println("Successfully connected to", dbType, "database!")
}
