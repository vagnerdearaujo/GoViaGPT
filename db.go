package db

import (
    "database/sql"
    "fmt"
    _ "github.com/denisenkom/go-mssqldb"
    _ "github.com/FirebirdSQL/firebird-go"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/lib/pq"
    "seuprojeto/config"
)

// Connect opens a connection to the specified database
func Connect() (*sql.DB, error) {
    dbConfig, err := config.LoadDBConfig()
    if err != nil {
        return nil, err
    }
    db, err := sql.Open(dbConfig.Type, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database))
    if err != nil {
        return nil, err
    }
    return db, nil
}
