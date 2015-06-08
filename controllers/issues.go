package controllers

import (
  "encoding/json"
  "net/http"
  "strconv"

  "github.com/julienschmidt/httprouter"

  "github.com/martin-brennan/hitch/data"
  "github.com/martin-brennan/hitch/errors"
)

var Issues = struct {
  Get func(w http.ResponseWriter, r *http.Request, p httprouter.Params)
  All func(w http.ResponseWriter, r *http.Request, p httprouter.Params)
}{
  Get: Get,
  All: All,
}

func Get(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
  id, err := strconv.Atoi(params.ByName("id"))
  if err != nil {
    hitchError.RaiseCustomError(w, params.ByName("id") + " is not a valid id", 400)
    return
  }

  issue, err := data.Issues.Get(id)
  if err != nil {
    hitchError.RaiseError(w, err)
    return
  }

  response, err := json.Marshal(issue)

  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}

func All(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  issues, err := data.Issues.All()
  if err != nil {
    hitchError.RaiseError(w, err)
    return
  }

  response, err := json.Marshal(issues)

  w.Header().Set("Content-Type", "application/json")
  w.Write(response)
}
