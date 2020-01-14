package main

import (
  "flag"
  "github.com/gookit/color"
  "github.com/traveltriangle/db-archiver/archive"
  "github.com/traveltriangle/db-archiver/config"
  "github.com/traveltriangle/db-archiver/query"
  "os"
)

func main(){
  parseFlags()
  columns, results, ids := query.Results()
  archive.ToCSV(columns, results)
  query.DeleteData(ids)
}

func parseFlags(){
  flag.StringVar(&config.Config.DbName, "db-name", "", "[REQUIRED] Name of Database")
  flag.StringVar(&config.Config.Table, "table", "", "[REQUIRED] table name to be archived")
  flag.StringVar(&config.Config.Where, "where", "",
    "condition to be used while archiving. Needed if --query is not provided")
  flag.StringVar(&config.Config.Query, "query", "",
    "if used it will ignore --where option. Needed if --where is not provided")
  flag.IntVar(&config.Config.Limit, "limit", 500, "limit the number of records")
  flag.IntVar(&config.Config.Batch, "batch", 0,
    "Fetch records in batch. If used it will ignore --limit")
  flag.StringVar(&config.Config.Path, "path", "/tmp/",
    "path to folder where the file will be stored")
  flag.StringVar(&config.Config.PrimaryKey, "pk", "id",
    "primary key which will be used to delete the records")
  flag.BoolVar(&config.Config.Optimize, "optimize", true, "Optimize Table after deletion")
  flag.BoolVar(&config.Config.Delete, "delete", true, "Delete from Table after archiving")
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