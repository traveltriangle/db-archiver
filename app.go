package main

import (
  "flag"
  "github.com/gookit/color"
  "github.com/traveltriangle/db-archiver/archive"
  "github.com/traveltriangle/db-archiver/config"
  "github.com/traveltriangle/db-archiver/query"
  "os"
  "time"
)

func main(){
  parseFlags()
  defer config.Config.Read.Db.Close()
  defer config.Config.Write.Db.Close()
  shouldStop := make(chan struct{})
  go func(){
    for {
      columns, results, ids := query.Results()
      if len(ids) == 0{
        shouldStop <- struct{}{}
      }
      archive.ToCSV(columns, results)
      query.DeleteData(ids)
      time.Sleep(time.Second * 10)
    }
  }()
  <-shouldStop
  query.OptimizeTable()
}

func parseFlags(){
  flag.StringVar(&config.Config.DbName, "db-name", "", "[REQUIRED] Name of Database")
  flag.StringVar(&config.Config.Table, "table", "", "[REQUIRED] table name to be archived")
  flag.StringVar(&config.Config.Where, "where", "",
    "condition to be used while archiving. Needed if --query is not provided")
  flag.StringVar(&config.Config.Query, "query", "",
    "if used it will ignore --where option. Needed if --where is not provided")
  flag.IntVar(&config.Config.Limit, "limit", 500, "limit the number of records")
  flag.StringVar(&config.Config.Path, "path", "/tmp/",
    "path to folder where the file will be stored")
  flag.StringVar(&config.Config.PrimaryKey, "pk", "id",
    "primary key which will be used to delete the records")
  flag.BoolVar(&config.Config.Optimize, "optimize", false, "Optimize Table after deletion")
  flag.Parse()
  if config.Config.DbName == "" {
    color.Error.Prompt("Database name is not specified")
    os.Exit(1)
  }
  if config.Config.Table == "" {
    color.Error.Prompt("Table name is not specified")
    os.Exit(1)
  }
  config.ConfigureMysql()
}