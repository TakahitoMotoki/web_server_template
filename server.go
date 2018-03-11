package main

import(
  "net/http"

  // MARK: - You MUST set your $GOPATH in your ~/zshrc
  "handlers"

  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()

  // MARK: - Routings
  router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
  router.HandleFunc("/test/{testId}", handlers.TestHandler).Methods("GET")

  http.Handle("/", router)
  http.ListenAndServe(":8080", nil)
}
