package cli

import (
	"fmt"
	"os"
)

func PWD() error {
	if dir, err := os.Getwd(); err != nil {
		return fmt.Errorf("pwd: %w", err)
	} else {
		fmt.Println(dir)
		return nil
	}

}
