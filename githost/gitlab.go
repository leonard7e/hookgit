package githost
import (
  "github.com/leonard7e/hookgit/appError"
  "github.com/leonard7e/hookgit/json"
)


func Host_Gitlab(h *GitHost) {
  h.HostName = "Gitlab"
  h.ObjectKind = func (data *json.JsonMap) (string) {
    str, ok := (*data)["object_kind"].(string)
    appError.AssertOk(ok, "object_kind nod defined in JSON")
    return str
  }
}
