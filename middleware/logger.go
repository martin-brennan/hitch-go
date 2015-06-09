package middleware

import (
  "log"
  "net/http"

  "github.com/julienschmidt/httprouter"

  "github.com/martin-brennan/hitch/errors"
)

func Logger(h httprouter.Handle) httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    log.Printf("[%s] %s", r.Method, r.URL.String())

    h(w, r, ps)
    return
  }
}

func Auth(h httprouter.Handle) httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

    if w.Header().Get("Authentication") == "" {
      log.Println("[400] no auth header")
      hitchError.RaiseCustomError(w, "no auth header", 401)
      return
    }

    log.Println("wooo")

    h(w, r, ps)
    return
  }
}

func HitchMiddleware(h httprouter.Handle) httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    Logger(Auth(h))(w, r, ps)
    return
  }
}

// middleware.HitchMiddleware(controllers.Issues.Get)
//
// HitchMiddleware -> controllers.Issues.Get
//
// h httprouter.Handle === controllers.Issues.Get
