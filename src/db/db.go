package db

import(
  "database/sql"

  // MARK: - MyPackage
  "models"

  _ "github.com/go-sql-driver/mysql"
)

func Index() []models.Person {
  db, err := sql.Open("mysql", "root:@/web_server_template")
  ShowError(err)
  defer db.Close()

  rows, err := db.Query("SELECT * FROM person")
  ShowError(err)

  // MARK: - Get columns name of params: rows
  // columns, err := rows.Columns()
  // ShowError(err)

  var people []models.Person
  person_tmp := models.Person{}
  for rows.Next() {
    err = rows.Scan(&person_tmp.Id, &person_tmp.Name, &person_tmp.Age)
    if err != nil {
      ShowError(err)
    }
    people = append(people, person_tmp)
  }

  return people
}

func ShowError(err error) {
  if err != nil {
    panic(err.Error())
  }
}
