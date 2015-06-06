package main

import (
  "database/sql"
  "time"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  "encoding/json"
)

type Issue struct {
  Id int
  Title string
  Description string
  Description_output string
  Created time.Time
  Modified time.Time
}

func main() {
  connection, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hitch?parseTime=true")
  defer connection.Close()

  if err != nil {
    panic(err)
  }

  row := connection.QueryRow("SELECT * FROM issues WHERE id=?", 1)
  issue := new(Issue)
  scanerr := row.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Description_output, &issue.Created, &issue.Modified)

  if scanerr != nil {
    panic(scanerr)
  }

  response, err := json.Marshal(issue)

  fmt.Println(string(response))
}
