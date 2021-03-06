package data

import (
  "database/sql"
  "time"

  // this is done because the compiler thinks it's not used,
  // when it is used in the call sql.Open("mysql")
  _ "github.com/go-sql-driver/mysql"

  "github.com/martin-brennan/hitch/config"
  "github.com/martin-brennan/hitch/models"
)

// Issues struct exposes the GetIssue and AllIssues
// methods for access like data.Issues.MethodName
var Issues = struct {
  Get func(int) (*models.Issue, error)
  All func() ([]*models.Issue, error)
  Add func(*models.Issue) (int64, error)
}{
  Get: GetIssue,
  All: AllIssues,
  Add: AddIssue,
}

func GetIssue(id int) (*models.Issue, error) {
  connection, err := sql.Open("mysql", config.Config["ConnectionString"])

  if err != nil {
    return nil, err
  }

  defer connection.Close()

  i := new(models.Issue)
  row := connection.QueryRow("SELECT * FROM issues WHERE id=?", id)
  scanerr := row.Scan(&i.Id, &i.Title, &i.Description, &i.Description_Output, &i.Created, &i.Modified)

  if scanerr != nil {
    return nil, scanerr
  }

  return i, nil
}

func AllIssues() ([]*models.Issue, error) {
  var issues []*models.Issue

  connection, err := sql.Open("mysql", config.Config["ConnectionString"])

  if err != nil {
    return nil, err
  }

  defer connection.Close()

  rows, err := connection.Query("SELECT * FROM issues")
  defer rows.Close()

  if err != nil {
    return nil, err
  }

  for rows.Next() {
    issue := new(models.Issue)
    err = rows.Scan(&issue.Id, &issue.Title, &issue.Description, &issue.Description_Output, &issue.Created, &issue.Modified)

    if err != nil {
      return nil, err
    }

    issues = append(issues, issue)
  }
  err = rows.Err()
  if err != nil {
    return nil, err
  }

  return issues, nil
}

func AddIssue(issue *models.Issue) (int64, error) {
  connection, err := sql.Open("mysql", config.Config["ConnectionString"])

  if err != nil {
    return 0, err
  }

  defer connection.Close()

  now := time.Now().UTC()

  result, err := connection.Exec(`INSERT INTO issues(title, description, description_output, created, modified) VALUES(?, ?, ?, ?, ?))`,
                          issue.Title, issue.Description, issue.Description_Output, now, now)

  if err != nil {
    return 0, err
  }

  id, err := result.LastInsertId()
  if err != nil {
    return 0, err
  }

  return id, nil
}
