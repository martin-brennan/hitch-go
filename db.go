package main

import (
  "encoding/json"
  "fmt"
  "github.com/martin-brennan/hitch/data"
)

func main() {

  issue := data.Issues.Get(1)
  issues := data.Issues.All()

  response, err := json.Marshal(issue)

  if err != nil {
    panic(err)
  }

  fmt.Println(string(response))

  response, err = json.Marshal(issues)

  if err != nil {
    panic(err)
  }

  fmt.Println(string(response))
}
