package main

import (
    "net/http"
    "github.com/leonard7e/hookgit/flags"
)
/*
"github.com/leonard7e/hookgit/callbacks"
"github.com/leonard7e/hookgit/githost"
*/


func main() {
  var pArgs = flags.ParseArgs()
  check_git_repository(pArgs)
  http.HandleFunc("/", CreateHandler(pArgs))
  http.ListenAndServe(":8080", nil)
}
