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
  Read struct{
    User       string `yaml:"user"`
    Password   string `yaml:"password"`
    Address    string `yaml:"address"`
    Db         *sql.DB
  } `yaml:"read"`
  Write struct{
    User       string `yaml:"user"`
    Password   string `yaml:"password"`
    Address    string `yaml:"address"`
    Db         *sql.DB
  } `yaml:"write"`
  DbName     string
  Table      string
  Where      string
  Query      string
  Limit      int
  Path       string
  PrimaryKey string
  Optimize   bool
}

var Config = conf{}

func ConfigureMysql() {
  dir, err := os.Getwd()
  HandleError(err, true)

  yamlFile, err := ioutil.ReadFile(filepath.Join(dir, "config", "db.yaml"))
  HandleError(err, true)
  err = yaml.Unmarshal(yamlFile, &Config)
  HandleError(err, true)

  configureReadDB()
  configureWriteDB()


}

func configureReadDB(){
  loc, err := time.LoadLocation("Asia/Kolkata")
  HandleError(err, true)
  mysqlCon := mysql.Config{
    User:                 Config.Read.User,
    Passwd:               Config.Read.Password,
    Addr:                 Config.Read.Address,
    Net:                  "tcp",
    DBName:               Config.DbName,
    Timeout:              0,
    ReadTimeout:          0,
    WriteTimeout:         0,
    Loc:                  loc,
    AllowNativePasswords: true,
  }

  Config.Read.Db, err = sql.Open("mysql", mysqlCon.FormatDSN())
  if err != nil {
    color.Error.Prompt(err.Error())
    panic(err)
  }

  if err := Config.Read.Db.Ping(); err != nil {
    color.Error.Prompt(err.Error())
    panic(err)
  }

}

func configureWriteDB(){
  loc, err := time.LoadLocation("Asia/Kolkata")
  HandleError(err, true)
  mysqlCon := mysql.Config{
    User:                 Config.Write.User,
    Passwd:               Config.Write.Password,
    Addr:                 Config.Write.Address,
    Net:                  "tcp",
    DBName:               Config.DbName,
    Timeout:              0,
    ReadTimeout:          0,
    WriteTimeout:         0,
    Loc:                  loc,
    AllowNativePasswords: true,
  }

  Config.Write.Db, err = sql.Open("mysql", mysqlCon.FormatDSN())
  if err != nil {
    color.Error.Prompt(err.Error())
    panic(err)
  }

  if err := Config.Write.Db.Ping(); err != nil {
    color.Error.Prompt(err.Error())
    panic(err)
  }
}
