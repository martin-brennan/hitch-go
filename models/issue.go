package models

import (
  "time"
)

type Issue struct {
  Id int `json:"id"`
  Title string `json:"title"`
  Description string `json:"description"`
  Description_Output string `json:"description_output"`
  Created time.Time `json:"created"`
  Modified time.Time `json:"modified"`
}
