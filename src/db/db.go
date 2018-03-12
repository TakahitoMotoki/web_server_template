package db

import(
  "database/sql"
  "fmt"

  _ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() {
  db, err := sql.Open("mysql", "root:@/web_server_template")
  if err != nil {
    fmt.Println("Database Connection Error!")
    panic(err.Error())
  }

  defer db.Close()
}


