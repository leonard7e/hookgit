package flags

import (
  "fmt"
  "github.com/leonard7e/hookgit/githost"
)

func chose_git_host (pArgs *PrgArg, host *string) {
  switch *host {
  case "gitlab":
    githost.Host_Gitlab(& pArgs.GitH)
  default:
    fmt.Println("Git host:", host)
    panic("Git Host not known")
  }
}
