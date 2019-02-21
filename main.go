package main

import (
	"bufio"
	"fmt"
	"os"
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

	fmt.Println("Select a branch number:")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		i := scanner.Text()

		if i == "Y" || i == "y" {
			break
		} else if i == "N" || i == "n" {
			result = false
			break
		} else {
			fmt.Println("yかnで答えてください。")
			fmt.Print(q)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	resultをゴニョゴニョ
}

func main() {
	showRemoteBranch()
}
