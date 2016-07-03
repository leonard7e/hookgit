package githost
import (
  "github.com/leonard7e/hookgit/json"
)


type GitHost struct {
  HostName string
  ObjectKind  func(*json.JsonMap) string
}

// Currently we only have a defined GitHost for Gitlab
