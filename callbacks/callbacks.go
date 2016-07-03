package callbacks

import (
  "github.com/leonard7e/hookgit/flags"
  "github.com/leonard7e/hookgit/json"
)

func RunCallbacks (args *flags.PrgArg, jM *json.JsonMap) {
    switch args.GitH.ObjectKind(jM) {
    case "push": // atm, we hardcode single callback
      git_pull()
    }
}
