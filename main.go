package main

import (
	"fm_tools/internal"
	"fmt"
)

func main() {
	exit := make(chan int, 1)
	for {
		select {
		case <-exit:
			return
		default:
			if err := internal.Run(&exit); err != nil {
				fmt.Println(err)
			}
		}
	}
}
