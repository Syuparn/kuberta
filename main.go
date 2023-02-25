package main

import (
	"fmt"
	"os"
)

func main() {
	// NOTE: os.Args[0] is command itself
	args := os.Args[1:]

	err, code := Exec(args, os.Stdout)

	if err != nil {
		fmt.Printf("kuberta: error: %v\n", err)
	}
	os.Exit(code)
}
