package main

import (
  "encoding/json"
  "fmt"
  "github.com/martin-brennan/hitch/models"
)

func main() {

  issue := new(models.Issue)
  issue.Get(1)

  response, err := json.Marshal(issue)

  if err != nil {
    panic(err)
  }

  fmt.Println(string(response))
}
