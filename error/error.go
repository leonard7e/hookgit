package error

import (
  "fmt"
)

func AssertNoErr (err error) {
  if err != nil {
    panic(err)
  }
}

func AssertOk (ok bool, ctx string) {
  if !ok {
    fmt.Println ("Assertion for", ctx, "went wrong!")
    panic("Assertion failed.")
  }
}
