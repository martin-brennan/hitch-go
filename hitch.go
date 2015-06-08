package main

import (
  "github.com/gorilla/mux"
  "net/http"

  "github.com/martin-brennan/hitch/controllers"
)

func main()  {
  router := mux.NewRouter()
  router.HandleFunc("/issue", controllers.Issues.All).Methods("GET")
  router.HandleFunc("/issue/{id}", controllers.Issues.Get).Methods("GET")

  http.ListenAndServe(":4556", router)
}
