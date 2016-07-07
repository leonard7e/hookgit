package flags

import (
  "fmt"
  "flag"
  "github.com/leonard7e/hookgit/githost"
  "github.com/leonard7e/hookgit/callbacks"
)

type PrgArg struct {
  GitH githost.GitHost
  Callbacks map[string]string
}

var (
  // f_host = flag.String("host", "gitlab", "Chose Git Host. One of these: (gitlab).")
  fa_cb FlagArray
)

var (
  pArgs PrgArg
)

func init_flags() {
  flag.Usage = func() {
    fmt.Println("Usage: hookgit [Flags]")
    fmt.Println("Flags are")
    flag.PrintDefaults()
  }

  flag.Var(&fa_cb, "cb", "Set Callback. Use syntax \"event.action\". E.G. hookgit -cb push.pull")
  pArgs.Callbacks = make(map[string]string)
}

func register_callbacks(cbs []string) {
  for i := range cbs {
    callbacks.RegisterCallback(cbs[i])
  }
}

func ParseArgs() *PrgArg {
  init_flags()
  flag.Parse()

  // Now we can use Options received from Flags
  if ( flag.NArg() >= 1 ) {
    flag.Usage()
    return nil
    // TODO: Return Error
  } else {
    register_callbacks(fa_cb)
    return &pArgs
  }
}
