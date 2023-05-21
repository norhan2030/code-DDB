package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main(){

dsn := "root:@tcp(localhost:3306)/ddb"

db, err := sql.Open("mysql", dsn)

// Check for errors when opening the connection
if err != nil {
    panic(err.Error())
}
err = db.Ping()

// Check for errors when pinging the database server
if err != nil {
    panic(err.Error())
}
//rows, err := db.Query("SELECT id, name FROM student")
rows, err := db.Query("SELECT id, name FROM std")

// Check for errors when executing the query
if err != nil {
    panic(err.Error())
}

defer rows.Close()

for rows.Next() {
    var id int
    var name string

    err := rows.Scan(&id, &name)

    // Check for errors when scanning the result set
    if err != nil {
        panic(err.Error())
    }

    fmt.Printf("id: %d, name: %s\n", id, name)
}
}