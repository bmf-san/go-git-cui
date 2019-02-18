package main

import (
	"flag"
	"fmt"
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
}
