package archive

import (
  "encoding/csv"
  "fmt"
  "github.com/gookit/color"
  "github.com/traveltriangle/db-archiver/config"
  "os"
  "time"
)

func ToCSV(columns []string, results []map[string]interface{},){
  color.Info.Prompt(fmt.Sprint( "Archiving ", len(results), " records."))
  now := time.Now()
  fileName := fmt.Sprintf("%s-%s-%d-%s-%d-%d:%d.csv", config.Config.DbName, config.Config.Table,
    now.Year(), now.Month().String(), now.Day(), now.Hour(), now.Minute())
  fileName = config.Config.Path + fileName

  color.Info.Prompt(fmt.Sprint( "Archiving to the file ", fileName))
  csvFile, err := os.Create( fileName)
  if err != nil {
    color.Error.Prompt(err.Error())
    os.Exit(1)
  }

  defer csvFile.Close()
  csvWriter := csv.NewWriter(csvFile)
  _ = csvWriter.Write(columns)
  for _, result := range results{
    var record =  make([]string, 0, 0)
    for _, col := range columns{
      record = append(record, fmt.Sprint(result[col]))
    }
    _ = csvWriter.Write(record)
  }
  csvWriter.Flush()
}