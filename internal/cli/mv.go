package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func MV(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("mv: expected 2 arguments, but %d were received", len(args))
	}

	info, err := os.Stat(args[0])
	if err != nil {
		return fmt.Errorf("mv: %w", err)
	}
	if info.IsDir() {
		if args[0] == args[1] {
			return nil
		}
		err := filepath.Walk(args[0], func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return fmt.Errorf("mv: %w", err)
			}

			rel, err := filepath.Rel(args[0], path)
			if err != nil {
				return fmt.Errorf("mv: относительный путь для %q: %w", path, err)
			}

			target := filepath.Join(args[1], rel)

			if info.IsDir() {
				return os.MkdirAll(target, 0755)
			}

			in, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("mv: открыть %q: %w", path, err)
			}
			defer in.Close()

			out, err := os.Create(target)
			if err != nil {
				return fmt.Errorf("mv: создать %q: %w", target, err)
			}
			defer out.Close()

			if _, err := io.Copy(out, in); err != nil {
				return fmt.Errorf("mv: копировать %q: %w", path, err)
			}

			return nil
		})
		if err != nil {
			return err
		}
		return os.RemoveAll(args[0])
	} else {
		if args[0] == args[1] {
			return nil
		}
		read, err := os.ReadFile(args[0])
		if err != nil {
			return fmt.Errorf("mv: %w", err)
		}
		if err := os.WriteFile(args[1], read, 0755); err != nil {
			return fmt.Errorf("mv: %w", err)
		}
		os.Remove(args[0])
	}
	return nil
}
