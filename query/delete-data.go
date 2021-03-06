package query

import (
  "context"
  "fmt"
  "github.com/gookit/color"
  "github.com/traveltriangle/db-archiver/config"
  "strings"
)

func DeleteData(ids []interface{}) {
  color.Info.Prompt(fmt.Sprint("Deleting from ", config.Config.Table))
  var idsString = make([]string, len(ids))
  for idx, id := range ids {
    idsString[idx] = id.(string)
  }
  query := fmt.Sprint("DELETE FROM ", config.Config.Table, " WHERE ID IN (", strings.Join(idsString, ","), ")")
  _, err := config.Config.Write.Db.ExecContext(context.Background(), query)
  config.HandleError(err, false)
  config.HandleError(err, true)
  color.Info.Prompt("Rows deleted %d from %s", len(ids), config.Config.Table)


}

func OptimizeTable(){
  if config.Config.Optimize {
    color.Info.Prompt(fmt.Sprint("Optimizing table ", config.Config.Table))
    query := fmt.Sprint("OPTIMIZE TABLE ", config.Config.Table)
    _, err := config.Config.Write.Db.ExecContext(context.Background(), query)
    config.HandleError(err, false)
  }
}
