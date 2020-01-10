package config

import "github.com/gookit/color"

func HandleError(err error, exit bool){
  if err != nil {
    color.Error.Prompt(err.Error())
    if exit{
      panic(err.Error())
    }
  }

}
