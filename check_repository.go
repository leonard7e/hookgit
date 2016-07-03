package main

import (
  "github.com/leonard7e/hookgit/flags"
  "github.com/leonard7e/hookgit/appError"
  "path"
  "io/ioutil"
  "fmt"
)



func check_git_repository(pArg *flags.PrgArg) {
  f, err := ioutil.ReadFile(path.Join(pArg.Repository,".git/config"))
  appError.AssertNoErr(err)
  // TODO ... Check Git for remote SSH (doesnt work with remote HTTP)

  fmt.Println(string(f))
}
