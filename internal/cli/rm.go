package cli

import (
	"fmt"
	"os"
)

func RM(args []string) error {
	if len(args) <= 0 {
		return fmt.Errorf("rm: No argument found")
	}
	if args[0] == "-d" {
		args = args[1:]
		if len(args) <= 0 {
			return fmt.Errorf("rm: No argument found")
		}
		for i := range args {
			if err := os.RemoveAll(args[i]); err != nil {
				return fmt.Errorf("rm: %w", err)
			}
		}
		return nil
	}
	for i := range args {
		if err := os.Remove(args[i]); err != nil {
			return fmt.Errorf("rm: %w", err)
		}
	}
	return nil
}
