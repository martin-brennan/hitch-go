package hitchError

import (
  "net/http"
  "encoding/json"
)

// ErrorData represents a JSON error message returned from the API
// for all 4-- and 5-- responses .
type ErrorData struct {
  Code int `json:"code"`
  Message string `json:"message"`
  ErrorMessage string `json:"error_message"`
}

// ErrorType is just a map of string representations of errors against
// ErrorData structs with friendly error messages e.g. sql: no rows is
// a 404 error with the message 'record not found'
var ErrorType = make(map[string]ErrorData)

// init sets up the ErrorType map entries
func init() {
  ErrorType["sql: no rows in result set"] = ErrorData{Code: 404, Message: "record not found", ErrorMessage: "sql: no rows in result set"}
}

// RaiseError responds with an ErrorData struct from the ErrorType
// map of string to ErrorData, based on the err message.
func RaiseError(w http.ResponseWriter, err error) {
  errorData := ErrorType[err.Error()]
  ErrorResponse(w, errorData)
}

// RaiseCustomError responds with an ErrorData struct with the provided
// message and error code as JSON.
func RaiseCustomError(w http.ResponseWriter, message string, code int) {
  errorData := ErrorData{Code: code, Message: message, ErrorMessage: ""}
  ErrorResponse(w, errorData)
}

// ErrorResponse uses an ErrorData struct, which is converted to JSON, to send error
// data back to the client.
func ErrorResponse(w http.ResponseWriter, err ErrorData) {
  responseData, jsonerr := json.Marshal(err)

  if jsonerr != nil {
    http.Error(w, "error processing error json", 500)
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(err.Code)
  w.Write(responseData)
}
