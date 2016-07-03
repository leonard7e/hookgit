package flags

import (
  "fmt"
  "flag"
  "github.com/leonard7e/hookgit/githost"
)
/*
"strings"
*/
// ---

type FlagArray []string

func (m *FlagArray) String() string {
	if len(*m) == 0 {
		return ""
	}
	return fmt.Sprint(*m)
}

func (m *FlagArray) Set(val string) error {
	(*m) = append(*m, val)
	return nil
}

// ---


type PrgArg struct {
  GitH githost.GitHost
  Repository string
  Callbacks map[string]string
}

func chose_git_host (pArgs *PrgArg, host *string) {
  switch *host {
  case "gitlab":
    githost.Host_Gitlab(& pArgs.GitH)
  default:
    fmt.Println("Git host:", host)
    panic("Git Host not known")
  }
}

// func register_callback(pArgs *PrgArg, cbstring *string) {
//   var cb = strings.Split(*cbstring,":")
//   pArgs.Callbacks[cb[0]] = cb[1]
// }

func ParseArgs() *PrgArg {
  flag.Usage = func() {
    fmt.Println("Usage: hookgit [Flags] Repository")
    fmt.Println("Flags are")
    flag.PrintDefaults()
  }

  // init Args
  pArgs := new(PrgArg)

  pArgs.Callbacks = make(map[string]string)
  gitHost := flag.String("host", "gitlab", "Get hooks from Gitlab (default)")

  var callbacks FlagArray
  flag.Var(&callbacks, "cb", "Set Callback")
 // parse flags
  flag.Parse()

  // Now we can use Options received from Flags
  if ( flag.NArg() == 1 ) {
    pArgs.Repository = flag.Arg(0)
    // register_callbacks(pArgs, callbacks)
    chose_git_host(pArgs, gitHost)
    return pArgs
  } else {
    flag.Usage()
    return nil // Not available in hookgit yet
  }
}
