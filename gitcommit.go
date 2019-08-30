package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	goPath, err := exec.Command("echo", os.Getenv("GOPATH")).Output()
	crPath := strings.Trim(string(goPath), "\n") + "/src/cr"
	gitPath := exec.Command("git", "commit -m ", "mani")
	gitPath.Dir = crPath
	res, err := gitPath.Output()
	fmt.Println(string(res), "Err ", err)

}
