package main
// "database/sql"
import (
    //"database/sql"
    "fmt"
    "net"
    _ "github.com/go-sql-driver/mysql"
)

func main(){

// dsn := "root:@tcp(localhost:3306)/testing"
// dsn := "root:@tcp(192.168.43.242:3306)/ddb"
db,err:=net.Dial("tcp","192.168.43.242:3306")
 if err!=nil{
  panic(err)
 }
 fmt.Printf("connect")
 defer db.Close()

//db, err := sql.Open("mysql", dsn)

// Check for errors when opening the connection
if err != nil {
    panic(err.Error())
}
//err = db.Ping()
// Check for errors when pinging the database server
if err != nil {
    panic(err.Error())
}
stmt, err := db.Query("SELECT id, name FROM student")
if err != nil {
    fmt.Println("Error preparing statement:", err.Error())
    return
}

defer stmt.Close()

// Execute the SELECT statement
rows, err := stmt.Query()
if err != nil {
    fmt.Println("Error executing query:", err.Error())
    return
}

defer db.Close()

// Process the results
for rows.Next() {
    var id int
    var name string
    err = rows.Scan(&id, &name)
    if err !=nil {
        fmt.Println("Error reading row:", err.Error())
        return
    }
    fmt.Printf("ID: %d, Name: %s\n", id, name)
}
}


// rows, err := db.Query("SELECT id, name FROM std")

// // Check for errors when executing the query
// if err != nil {
//     panic(err.Error())
// }
// defer rows.Close()

// for rows.Next() {
//     var id int
//     var name string
//     err := rows.Scan(&id, &name)
//     // Check for errors when scanning the result set
//     if err != nil {
//         panic(err.Error())
//     }
//     fmt.Printf("id: %d, name: %s\n", id, name)
// }

// }