package main

import (
  "fmt"
  "net/http"
  "github.com/leonard7e/hookgit/flags"
  "github.com/leonard7e/hookgit/callbacks"
  "github.com/leonard7e/hookgit/json"
)

func CreateHandler(args *flags.PrgArg) (func(w http.ResponseWriter, r *http.Request)) {
  return func (w http.ResponseWriter, r *http.Request) {
    fmt.Println("HTTP-Request came from host", r.Host)
    if (r.Method == "POST") {
      fmt.Println("Post method received")
      callbacks.RunCallbacks(& args.GitH, json.ReadJson(r))
    } else {
      fmt.Println("Method:", r.Method)
    }

  }
}
