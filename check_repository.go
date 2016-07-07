package main

import (
  "github.com/leonard7e/hookgit/flags"
  "github.com/leonard7e/hookgit/error"
  "github.com/leonard7e/hookgit/githost"
  "io/ioutil"
  "fmt"
  "regexp"
)

func check_git_repository(pArg *flags.PrgArg) {
  f, err := ioutil.ReadFile(".git/config")
  error.AssertNoErr(err)
  // TODO ... Check Git for remote SSH (doesnt work with remote HTTP)

  var re_url = regexp.MustCompile("url = .*\n")
  var re_ssh = regexp.MustCompile("git@.*")
  var re_host = regexp.MustCompile("(github)|(gitlab)")

  var url = re_url.Find(f)

  if re_ssh.Match(url) {
    fmt.Println("Ok. This repository uses Ssh for Push/Pull.")
    // Chose Githost service (Gitlab / ... )
    var hostStr = string(re_host.Find(url))
    githost.ChoseGithost(&pArg.GitH, &hostStr)
  } else {
    panic("Please use SSH based clones.")
  }
}
