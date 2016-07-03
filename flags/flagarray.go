package flags

import "fmt"

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
