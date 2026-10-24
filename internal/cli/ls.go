package cli

import (
	"fmt"
	"os"
)

func LS(args []string) error {
	if len(args) <= 0 {
		if output, err := os.ReadDir("./"); err != nil {
			return fmt.Errorf("ls: %w", err)
		} else {
			for i := range output {
				if output[i].IsDir() {
					fmt.Printf("%s/ ", output[i].Name())
				} else {
					fmt.Printf("%s ", output[i].Name())
				}
			}
			fmt.Println()
		}
	}
	for i := range args {
		if output, err := os.ReadDir(args[i]); err != nil {
			return fmt.Errorf("ls: %w", err)
		} else {
			for i := range output {
				if output[i].IsDir() {
					fmt.Printf("%s/ ", output[i].Name())
				} else {
					fmt.Printf("%s ", output[i].Name())
				}
			}
			fmt.Println()
		}
	}
	return nil
}
