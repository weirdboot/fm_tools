package cli

import (
	"fmt"
	"os"
)

func CD(args []string) error {
	if len(args) <= 0 {
		return fmt.Errorf("cd: No argument found")
	} else if len(args) > 1 {
		return fmt.Errorf("cd: expected 1 argument, but %d were received", len(args))
	}
	return os.Chdir(args[0])
}
