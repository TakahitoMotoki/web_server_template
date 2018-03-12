package db

import(
  "database/sql"
  "fmt"

  _ "github.com/go-sql-driver/mysql"
)

/*
func ConnectToDB() {
  db, err := sql.Open("mysql", "root:@/web_server_template")
  if err != nil {
    fmt.Println("Database Connection Error!")
    panic(err.Error())
  }
}
*/

func Index() {
  var table = "person"

  db, err := sql.Open("mysql", "root:@/web_server_template")
  ShowError(err)
  defer db.Close()

  rows, err := db.Query("SELECT * FROM person")
  ShowError(err)

  // MARK: - Get columns name of params: rows
  columns, err := rows.Columns()
  ShowError(err)

  results, err := db.Prepare(fmt.Sprintf("SELECT * FROM %s", table))
  ShowError(err)
  defer results.Close()

  fmt.Println(results)
  fmt.Println("+=+=+=+=+=+=+=+=+=+=+")
  fmt.Println(columns)
}

func ShowError(err error) {
  if err != nil {
    panic(err.Error())
  }
}
