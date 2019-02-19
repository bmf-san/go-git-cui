package main

import (
	"fmt"
	"os/exec"
)

func gitCheckoutToRemoteBranch() {
	out, err := exec.Command("git", "branch", "-r").Output()

	if err != nil {
		panic(err)
	}

	fmt.Println("%T", out)

	// scanner := bufio.NewScanner(os.Stdin)
	// for scanner.Scan() {
	// 	i := scanner.Text()

	// 	if i == "Y" || i == "y" {
	// 		break
	// 	} else if i == "N" || i == "n" {
	// 		result = false
	// 		break

	// 		fmt.Println("yかnで答えてください。")
	// 		fmt.Print(q)
	// 	}
	// }

	// if err := scanner.Err(); err != nil {
	// 	panic(err)
	// }

	// if result {

	// } else {

	// }

	// fmt.Println(string(out))
}

func main() {
	gitCheckoutToRemoteBranch()
}
