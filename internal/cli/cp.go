package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func CP(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("cp: expected 2 arguments, but %d were received", len(args))
	}

	info, err := os.Stat(args[0])
	if err != nil {
		return fmt.Errorf("cp: %w", err)
	}
	if info.IsDir() {
		return filepath.Walk(args[0], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("cp: %w", err)
			}
			rel, _ := filepath.Rel(args[0], path)
			target := filepath.Join(args[1], rel)

			if info.IsDir() {
				return os.MkdirAll(target, 0755)
			}

			in, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("cp: %w", err)
			}
			defer in.Close()

			out, err := os.Create(target)
			if err != nil {
				return fmt.Errorf("cp: %w", err)
			}
			defer out.Close()

			_, err = io.Copy(out, in)
			if err != nil {
				return fmt.Errorf("cp: %w", err)
			}
			return nil
		})
	} else {
		read, err := os.ReadFile(args[0])
		if err != nil {
			return fmt.Errorf("cp: %w", err)
		}
		if err := os.WriteFile(args[1], read, 0755); err != nil {
			return fmt.Errorf("cp: %w", err)
		}
	}

	return nil
}
