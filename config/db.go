package config

import (
  "database/sql"
  "github.com/go-sql-driver/mysql"
  "github.com/gookit/color"
  "time"
)

var Config = struct {
  User       string
  Password   string
  Address    string
  DbName     string
  Table      string
  Where      string
  Query      string
  Limit      int
  Batch      int
  Path       string
  PrimaryKey string
  Db         *sql.DB
  Optimize   bool
  Delete     bool
}{}

func ConfigureMysql() {
  loc, err := time.LoadLocation("Asia/Kolkata")

  if err != nil {
    color.Error.Prompt(err.Error())
    panic(err)
  }

  mysqlCon := mysql.Config{
    User:                 Config.User,
    Passwd:               Config.Password,
    Net:                  "tcp",
    Addr:                 Config.Address,
    DBName:               Config.DbName,
    Timeout:              0,
    ReadTimeout:          0,
    WriteTimeout:         0,
    Loc:                  loc,
    AllowNativePasswords: true,
  }
  Config.Db, err = sql.Open("mysql", mysqlCon.FormatDSN())
  if err != nil {
    color.Error.Prompt(err.Error())
    panic(err)
  }

  if err := Config.Db.Ping(); err != nil {
    color.Error.Prompt(err.Error())
    panic(err)
  }
}
