package main

import (
    "database/sql"
    "fmt"
    _ "github.com/denisenkom/go-mssqldb"
)

type DBExample struct {
    ConnectionUrl string
    Db            *sql.DB
    Stmt          *sql.Stmt
    Rows          *sql.Rows
}

func (d *DBExample) SetConnectionUrl(connectionUrl string) {
    d.ConnectionUrl = connectionUrl
}

func (d *DBExample) GetConnectionUrl() string {
    return d.ConnectionUrl
}

func (d *DBExample) GetAllClients() {
    var err error
    d.Db, err = sql.Open("mssql", d.ConnectionUrl)
    if err != nil {
        panic(err)
    }
    defer d.Db.Close()

    d.Stmt, err = d.Db.Prepare("SELECT id, nome, codigo_interno FROM clientes")
    if err != nil {
        panic(err)
    }
    defer d.Stmt.Close()

    d.Rows, err = d.Stmt.Query()
    if err != nil {
        panic(err)
    }
    defer d.Rows.Close()
}

func (d *DBExample) Close() {
    d.Db.Close()
    d.Stmt.Close()
    d.Rows.Close()
}

func main() {
    db := &DBExample{}
    db.SetConnectionUrl("server=hostname;user id=username;password=password;port=port;database=dbname")
    db.GetAllClients()

    for db.Rows.Next() {
        var id int
        var nome string
        var codigoInterno int
        db.Rows.Scan(&id, &nome, &codigoInterno)
        fmt.Println(id, nome, codigoInterno)
    }

    db.Close()
}
