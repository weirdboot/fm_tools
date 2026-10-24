package internal

import (
	"bufio"
	"fm_tools/internal/cli"
	"fmt"
	"io"
	"os"
	"strings"
)

func Run(exit *chan int) error {
	if dir, err := os.Getwd(); err != nil {
		return err
	} else {
		fmt.Printf("fm [%s]> ", dir)
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return err
	}

	if len(input) <= 0 || input == "\n" {
		return fmt.Errorf("fm_tools: Unknown command! (Type `help` for help!)")
	}

	args := parseInput(input)
	switch args[0] {
	case "cd":
		return cli.CD(args[1:])
	case "cp":
		return cli.CP(args[1:])
	case "grep":
		return cli.GREP(args[1:])
	case "ls":
		return cli.LS(args[1:])
	case "mv":
		return cli.MV(args[1:])
	case "rm":
		return cli.RM(args[1:])
	case "pwd":
		return cli.PWD()
	case "help":
		return cli.HELP()
	case "exit":
		*exit <- 1
		return nil
	case "clear":
		fmt.Print("\033[H\033[2J")
		return nil
	default:
		return fmt.Errorf("fm_tools: Unknown command! (Type `help` for help!)")
	}
}

func parseInput(input string) []string {
	parts := strings.Fields(strings.ReplaceAll(input, `\ `, "\x00"))
	for i := range parts {
		parts[i] = strings.ReplaceAll(parts[i], "\x00", " ")
	}
	return parts
}
