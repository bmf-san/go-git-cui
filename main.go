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

	fmt.Println(string(out))

	// 文字列を1行入力
	// func StrStdin() (stringInput string) {
	// 	scanner := bufio.NewScanner(os.Stdin)

	// 	scanner.Scan()
	// 	stringInput = scanner.Text()

	// 	stringInput = strings.TrimSpace(stringInput)
	// 	return
	// }

	// func main() {
	// 	p := StrStdin()
	// 	fmt.Println(p)
	// }
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
