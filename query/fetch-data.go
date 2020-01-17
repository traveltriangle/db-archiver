package query

import (
  "database/sql"
  "fmt"
  "github.com/gookit/color"
  "github.com/traveltriangle/db-archiver/config"
  "os"
)

func Results() ([]string, []map[string]interface{}, []interface{}){
  var results = make([]map[string]interface{},0, 0)
  var columns = make([]string, 0, 0)
  var ids = make([]interface{},0, 0)
  var query string
  if config.Config.Query != ""{
    query = config.Config.Query

  } else if config.Config.Where != ""{
    query = fmt.Sprint("SELECT SQL_NO_CACHE * FROM ", config.Config.Table, " WHERE ",
      config.Config.Where)
  } else {
    color.Error.Prompt("Any one of --query or --where should be specified.")
    os.Exit(1)
  }
  columns = fetchData(config.Config.Limit, query, &results, &ids)

 return columns, results, ids

}

func fetchData(limit int, query string, results *[]map[string]interface{}, ids *[]interface{}) ([]string){
  var rows *sql.Rows
  var err error
  numOfRows := 0
  query = fmt.Sprint(query, " LIMIT ", limit)
  color.Info.Prompt(fmt.Sprint("Running query: ", query))
  rows, err = config.Config.Read.Db.Query(query)
  config.HandleError(err, true)
  defer rows.Close()

  columns, err := rows.Columns()
  config.HandleError(err, true)

  values := make([]sql.RawBytes, len(columns))
  scanArgs := make([]interface{}, len(values))
  for i := range values {
    scanArgs[i] = &values[i]
  }
  for rows.Next() {
    numOfRows += 1
    err = rows.Scan(scanArgs...)
    config.HandleError(err, true)

    var value interface{}
    var result = make(map[string]interface{})
    for i, col := range values {
      if col == nil {
        value = "NULL"
      } else {
        value = string(col)
      }
      result[columns[i]] = value
    }
    if result[config.Config.PrimaryKey] == ""{
      color.Error.Prompt(fmt.Sprint("Primary Key ", config.Config.PrimaryKey,
        " not present. Cancelling the Archival."))
    }
    *ids = append(*ids, result[config.Config.PrimaryKey])
    *results = append(*results, result)
  }
  config.HandleError(rows.Err(), true)

  if numOfRows == 0{
    return columns
  }
  return columns
}
