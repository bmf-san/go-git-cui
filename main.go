package main

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

// 最終的にバイナリ配布を想定して、コマンド関数はオプション（エイリアス名）またはサブコマンド（ex. go-git g-c-r-b)
func showRemoteBranch() {
	out, err := exec.Command("git", "branch", "-r").Output()

	if err != nil {
		panic(err)
	}

	var branches []string

	for _, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(out), -1) {
		if v != "" {
			branches = append(branches, strings.TrimSpace(string(v)))
		}
	}

	fmt.Println("Remote Branches:")
	for i, v := range branches {
		fmt.Println(i, v)
	}
}

func main() {
	showRemoteBranch()
}
