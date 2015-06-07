package main

import (
  "github.com/codegangsta/negroni"
  "net/http"
  "encoding/json"
)

type Hello struct {
  Id int
  Name string
}

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    hi := Hello{Id: 1, Name: "Martin"}

    response, err := json.Marshal(hi)

    if err != nil {
      panic(err)
    }

    w.Header().Set("Content-Type", "application/json")

    w.Write([]byte(response))
  })

  server := negroni.Classic()
  server.UseHandler(mux)
  server.Run(":4444")
}
