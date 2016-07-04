package main

import (
  "github.com/leonard7e/hookgit/flags"
  "github.com/leonard7e/hookgit/error"
  "path"
  "io/ioutil"
  "fmt"
  "regexp"
)



func check_git_repository(pArg *flags.PrgArg) {
  f, err := ioutil.ReadFile(path.Join(pArg.Repository,".git/config"))
  error.AssertNoErr(err)
  // TODO ... Check Git for remote SSH (doesnt work with remote HTTP)
  var sshGit = regexp.MustCompile("url = git@.*\\.com.*")
  if sshGit.Match(f) {
    fmt.Println("Ok. This repository uses Ssh for Push/Pull.")
  } else {
    panic("Please use SSH based clones.")
  }
}
