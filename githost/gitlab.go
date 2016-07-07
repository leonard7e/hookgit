package githost
import (
  "github.com/leonard7e/hookgit/error"
  "github.com/leonard7e/hookgit/json"
)


func HostGitlab(h *GitHost) {
  h.HostName = "Gitlab"
  h.ObjectKind = func (data *json.JsonMap) (string) {
    str, ok := (*data)["object_kind"].(string)
    error.AssertOk(ok, "object_kind nod defined in JSON")
    return str
  }
}
