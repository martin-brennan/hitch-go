package main

import (
  "database/sql"
  "time"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
)

type Issue struct {
  id int
  title string
  description string
  description_output string
  created time.Time
  modified time.Time
}

func main() {
  connection, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/hitch?parseTime=true")
  defer connection.Close()

  if err != nil {
    panic(err)
  }

  row := connection.QueryRow("SELECT * FROM issues WHERE id=?", 1)
  issue := new(Issue)
  scanerr := row.Scan(&issue.id, &issue.title, &issue.description, &issue.description_output, &issue.created, &issue.modified)

  if scanerr != nil {
    panic(scanerr)
  }

  fmt.Printf(issue.title + "\n")
}
