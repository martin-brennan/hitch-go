package middleware

import (
  "fmt"
  "net/http"
)
// 
// func Logger(next func) http.Handler {
//   return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
//     fmt.Println(r.Method + " - " + r.URL.String())
//
//     next.ServeHTTP(w, r)
//   })
// }
//
// func Logger(next func)
