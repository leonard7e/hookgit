package callbacks

import (
  "github.com/leonard7e/hookgit/githost"
  "github.com/leonard7e/hookgit/json"
  "strings"
  "fmt"
)

var cbList map[string][]string

func InitCallbacks() {
  cbList = make(map[string][]string)
}

// TODO: We need to check for dangerous combinations. E.G. -cb push.push would be problematic
func RunCallbacks (gitH *githost.GitHost, jM *json.JsonMap) {
  var kind = gitH.ObjectKind(jM)
  var actions = cbList[kind]
  for i := range actions {
    switch actions[i] {
    case "pull":
      git_pull ()
    default:
      fmt.Println("Callback", actions[i], "not supported for event", kind)
    }
  }
}

// site Effect. Needs Error. Will do later.
func RegisterCallback (cb string) {
  var cbs = strings.Split(cb,".")
  fmt.Println("Registering Callbacks.", cbs)
  cbList[cbs[0]] = append(cbList[cbs[0]], cbs[1])
}
