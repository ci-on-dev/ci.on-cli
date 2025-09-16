package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ci.on <init|test> [options]")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "init":
		if err := RunInit(); err != nil {
			fmt.Println("Init failed:", err)
			os.Exit(1)
		}
	case "test":
		RunTest(os.Args[2:])
	default:
		fmt.Println("Unknown command:", os.Args[1])
		os.Exit(1)
	}

	os.RemoveAll("tmp")
}
