package handlers

import(
  //"encoding/json"
  "fmt"
  //"io"
  //"io/ioutil"
  "net/http"
  //"strconv"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Welcome!")
}

func TestHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Println("Test")
}
