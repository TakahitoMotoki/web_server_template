package handlers

import(
  "encoding/json"
  "fmt"
  // "io"
  // "io/ioutil"
  "net/http"
  // "strconv"

  // MARK: - My Package
  "models"
  "db"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Welcome!")
  person1 := models.Person{Id: 1, Name: "Takahito Motoki", Age: 21}

  if err := json.NewEncoder(w).Encode(person1); err != nil {
    panic(err)
  }

  db.Index()
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
  person1 := models.Person{Id: 1, Name: "Takahito Motoki", Age: 21}
  fmt.Printf("%d", person1.GetId())
}
