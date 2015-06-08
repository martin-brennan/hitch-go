package controllers

import (
  "encoding/json"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"

  "github.com/martin-brennan/hitch/data"
  "github.com/martin-brennan/hitch/errors"
)

var Issues = struct {
  Get func(w http.ResponseWriter, r *http.Request)
  All func(w http.ResponseWriter, r *http.Request)
}{
  Get: Get,
  All: All,
}

func Get(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
  id, err := strconv.Atoi(params["id"])
  if err != nil {
    hitch_error.RaiseCustomError(w, params["id"] + " is not a valid id", 400)
    return
  }

  issue, err := data.Issues.Get(id)
  if err != nil {
    hitch_error.RaiseError(w, err)
    return
  }

  response, err := json.Marshal(issue)

  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func All(w http.ResponseWriter, r *http.Request) {
  issues, err := data.Issues.All()
  if err != nil {
    hitch_error.RaiseError(w, err)
    return
  }

  response, err := json.Marshal(issues)

  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}
