package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var (
		msg   = flag.String("m", "default message.", "Message")
		num   = flag.Uint("n", 0, "Count(>=0)")
		debug = flag.Bool("debug", false, "Debug Mode enabled?")
	)

	flag.Parse()

	fmt.Printf("param -m -> %s\n", *msg) // as -m option
	fmt.Printf("param -n -> %s\n", *num) // as -n option
	fmt.Printf("param -debug -> %t\n", *debug)

	if Question("今日は元気ですか？[y/n] ") {
		out, err := exec.Command("ls").Output()

		if err != nil {
			panic(err)
		}

		fmt.Printf("ls result: \n%s", string(out))
	} else {
		fmt.Println("元気出してくださいね！")
	}
}

func Question(q string) bool {
	result := true
	fmt.Print(q)

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
	return result
}
