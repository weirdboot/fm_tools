package main

import (
	"fm_tools/internal"
	"fmt"
)

func main() {
	for {
		shouldExit, err := internal.Run()
		if err != nil {
			fmt.Println(err)
		}
		if shouldExit {
			return
		}
	}
}
