package config

import (
  "database/sql"
  "github.com/go-sql-driver/mysql"
  "github.com/gookit/color"
  "gopkg.in/yaml.v2"
  "io/ioutil"
  "os"
  "path/filepath"
  "time"
)

type conf struct {
  User       string `yaml:"user"`
  Password   string `yaml:"password"`
  Address    string `yaml:"address"`
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
}

var Config = conf{}

func ConfigureMysql() {
  dir, err := os.Getwd()
  HandleError(err, true)

  loc, err := time.LoadLocation("Asia/Kolkata")
  HandleError(err, true)

  yamlFile, err := ioutil.ReadFile(filepath.Join(dir, "config", "db.yaml"))
  HandleError(err, true)
  err = yaml.Unmarshal(yamlFile, &Config)
  HandleError(err, true)

  mysqlCon := mysql.Config{
    User:                 Config.User,
    Passwd:               Config.Password,
    Addr:                 Config.Address,
    Net:                  "tcp",
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
