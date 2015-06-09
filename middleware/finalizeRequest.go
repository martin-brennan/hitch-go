package middleware

import (
  "log"
  "net/http"
)

func FinalizeRequest(w http.ResponseWriter, request *http.Request, response []byte, contentType string, code int) {
  log.Printf("[%s(%d)] %s", request.Method, code, request.URL.String())

  w.Header().Set("Content-Type", contentType)
  w.WriteHeader(code)
  w.Write(response)
}
