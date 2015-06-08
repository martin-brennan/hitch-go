package hitch_error

import (
  "net/http"
  "encoding/json"
  "fmt"
)

type ErrorData struct {
  Code int `json:"code"`
  Message string `json:"message"`
  ErrorMessage string `json:"error_message"`
}

var ErrorType = make(map[string]ErrorData)

func init() {
  ErrorType["sql: no rows in result set"] = ErrorData{Code: 404, Message: "record not found", ErrorMessage: "sql: no rows in result set"}
}

func RaiseError(w http.ResponseWriter, err error) {
  fmt.Println(ErrorType)
  errorData := ErrorType[err.Error()]
  ErrorResponse(w, errorData)
}

func RaiseCustomError(w http.ResponseWriter, message string, code int) {
  errorData := ErrorData{Code: code, Message: message, ErrorMessage: ""}
  ErrorResponse(w, errorData)
}

func ErrorResponse(w http.ResponseWriter, err ErrorData) {
  responseData, jsonerr := json.Marshal(err)

  if jsonerr != nil {
    http.Error(w, "error processing error json", 500)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(err.Code)
  w.Write(responseData)
}
