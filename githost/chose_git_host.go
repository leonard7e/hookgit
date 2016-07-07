package githost

import (
  "fmt"
)

func ChoseGithost (gith *GitHost, host *string) {
  switch *host {
  case "gitlab":
    HostGitlab(gith)
  default:
    fmt.Println("Git host:", host)
    panic("Git Host not known")
  }
}
