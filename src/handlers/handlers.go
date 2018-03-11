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
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Welcome!")
  p1 := models.Person{Id: 1, Name: "Takahito Motoki", Age: 21}
  // p2 := models.Person{Id: 2, Name: "Takao Baba", Age: 21}

  if err := json.NewEncoder(w).Encode(p1); err != nil {
    panic(err)
  }
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Test")
}
