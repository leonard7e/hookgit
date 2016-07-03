package flags

import (
  "fmt"
  "flag"
  "github.com/leonard7e/hookgit/githost"
)

type PrgArg struct {
  GitH githost.GitHost
  Repository string
  Callbacks map[string]string
}

var (
  gitHost = flag.String("host", "gitlab", "Get hooks from Gitlab (default)")
)

var (
  callbacks FlagArray
)

var (
  pArgs PrgArg
)

func init_flags() {
  flag.Usage = func() {
    fmt.Println("Usage: hookgit [Flags] Repository")
    fmt.Println("Flags are")
    flag.PrintDefaults()
  }

  flag.Var(&callbacks, "cb", "Set Callback")
  pArgs.Callbacks = make(map[string]string)
}

func ParseArgs() *PrgArg {
  init_flags()
  flag.Parse()

  // Now we can use Options received from Flags
  if ( flag.NArg() == 1 ) {
    pArgs.Repository = flag.Arg(0)
    // register_callbacks(pArgs, callbacks)
    chose_git_host(&pArgs, gitHost)
    return &pArgs
  } else {
    flag.Usage()
    return nil // Not available in hookgit yet
  }
}
