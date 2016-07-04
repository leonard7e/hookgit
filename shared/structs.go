package shared

type PrgArg struct {
  GitH githost.GitHost
  Repository string
  Callbacks map[string]string
}

type FlagArray []string
