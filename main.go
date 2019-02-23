package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func getRemoteBranches() (branches []string) {
	out, err := exec.Command("git", "branch", "-r").Output()

	if err != nil {
		panic(err)
	}

	for _, v := range regexp.MustCompile("\r\n|\n\r|\n|\r").Split(string(out), -1) {
		if v != "" {
			branches = append(branches, strings.TrimSpace(string(v)))
		}
	}

	return branches
}

func checkOutToRemoteBranch(branches []string) {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())

	_, err := exec.Command("git", "checkout", "-b", branches[input], branches[input]).Output()

	if err != nil {
		panic(err)
	}
}

func main() {
	branches := getRemoteBranches()
	fmt.Println("Remote Branches:")
	for i, v := range branches {
		fmt.Println(i, v)
	}

	fmt.Println("Select a branch number:")
	checkOutToRemoteBranch(branches)
}
