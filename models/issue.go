package models

import (
  "time"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "github.com/martin-brennan/hitch/config"
)

type Issue struct {
  Id int `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
  Description_Output string `json:"description_output"`
  Created time.Time `json:"created"`
  Modified time.Time `json:"modified"`
}

func (i *Issue) Get(id int) {
  connection, err := sql.Open("mysql", config.Config.ConnectionString)
  defer connection.Close()

  if err != nil {
    panic(err)
  }

  row := connection.QueryRow("SELECT * FROM issues WHERE id=?", id)
  scanerr := row.Scan(&i.Id, &i.Title, &i.Description, &i.Description_Output, &i.Created, &i.Modified)

  if scanerr != nil {
    panic(scanerr)
  }
}
