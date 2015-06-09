package main

import (
  "log"
  "net/http"

  "github.com/julienschmidt/httprouter"

  "github.com/martin-brennan/hitch/controllers"
  "github.com/martin-brennan/hitch/middleware"
)

func main()  {
  router := httprouter.New()
  router.GET("/issue", middleware.Logger(controllers.Issues.All))
  router.GET("/issue/:id", controllers.Issues.Get)

  log.Fatal(http.ListenAndServe(":4556", router))
}
