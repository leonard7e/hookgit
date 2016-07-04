package callbacks

import (
  "os/exec"
  "fmt"
  "github.com/leonard7e/hookgit/error"
)

/*
"github.com/leonard7e/hookgit/flags"
"path"
"io/ioutil"
*/
func cmd_git (cmd string) {
  git_cmd := exec.Command("git", cmd)
  o, err := git_cmd.Output()
  error.AssertNoErr(err)

  git_cmd.Run()
  fmt.Println(string(o))
}

func git_pull () {
  cmd_git("pull")
}
