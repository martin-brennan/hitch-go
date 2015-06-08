package main

import (
  "log"
  "net/http"

  "github.com/julienschmidt/httprouter"

  "github.com/martin-brennan/hitch/controllers"
)

func main()  {
  router := httprouter.New()
  router.GET("/issue", controllers.Issues.All)
  router.GET("/issue/:id", controllers.Issues.Get)

  log.Fatal(http.ListenAndServe(":4556", router))
}
