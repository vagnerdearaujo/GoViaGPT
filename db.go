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
    var connString string
    switch dbConfig.Type {
    case "sqlserver":
        connString = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;", dbConfig.Host, dbConfig.User, dbConfig.Password, dbConfig.Port, dbConfig.Database)
    case "firebird":
        connString = fmt.Sprintf("sysdba:masterkey@%s:%s/%s", dbConfig.Host, dbConfig.Port, dbConfig.Database)
    case "mysql":
        connString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Port, dbConfig.Database)
    case "postgres":
        connString = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Database)
    }
    db, err := sql.Open(dbConfig.Type, connString)
    if err != nil {
        return nil, err
    }
    return db, nil
}
