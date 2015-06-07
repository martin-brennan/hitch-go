package data

import (
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/martin-brennan/hitch/config"
  "github.com/martin-brennan/hitch/models"
)

var Issues = struct {
  Get func(int) (*models.Issue)
  All func() ([]*models.Issue)
}{
  Get: GetIssue,
  All: AllIssues,
}

func GetIssue(id int) (*models.Issue) {
  connection, err := sql.Open("mysql", config.Config.ConnectionString)
  defer connection.Close()

  if err != nil {
    panic(err)
  }

  i := new(models.Issue)
  row := connection.QueryRow("SELECT * FROM issues WHERE id=?", id)
  scanerr := row.Scan(&i.Id, &i.Title, &i.Description, &i.Description_Output, &i.Created, &i.Modified)

  if scanerr != nil {
    panic(scanerr)
  }

  return i
}

func AllIssues() ([]*models.Issue) {
  var issues []*models.Issue

  connection, err := sql.Open("mysql", config.Config.ConnectionString)
  defer connection.Close()

  if err != nil {
    panic(err)
  }

  rows, err := connection.Query("SELECT * FROM issues")
  defer rows.Close()

  if err != nil {
    panic(err)
  }

  for rows.Next() {
    issue := new(models.Issue)
    err = rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Description_Output, &issue.Created, &issue.Modified)

    if err != nil {
      panic(err)
    }

    issues = append(issues, issue)
  }
  err = rows.Err()

  return issues
}
