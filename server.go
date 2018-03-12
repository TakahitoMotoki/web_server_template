package main

import(
  "net/http"

  // MARK: - You MUST set your $GOPATH in your ~/zshrc
  "handlers"

  "github.com/gorilla/mux"
)

func main() {
  router := mux.NewRouter()

  // MARK: - Routing Zone Start
  router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
  router.HandleFunc("/test/{testId}", handlers.TestHandler).Methods("GET")
  // MARK: - Routing Zone End

  http.Handle("/", router)
  http.ListenAndServe(":8080", nil)
}
