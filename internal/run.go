package internal

import (
	"bufio"
	"fm_tools/internal/cli"
	"fmt"
	"io"
	"os"
	"strings"
)

func Run() (bool, error) {
	if dir, err := os.Getwd(); err != nil {
		return false, err
	} else {
		fmt.Printf("fm [%s]> ", dir)
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return false, err
	}
	if err == io.EOF {
		return true, nil
	}

	if len(input) <= 0 || input == "\n" {
		return false, fmt.Errorf("fm_tools: Unknown command! (Type `help` for help!)")
	}

	args := parseInput(input)
	switch args[0] {
	case "cd":
		return false, cli.CD(args[1:])
	case "cp":
		return false, cli.CP(args[1:])
	case "grep":
		return false, cli.GREP(args[1:])
	case "ls":
		return false, cli.LS(args[1:])
	case "mv":
		return false, cli.MV(args[1:])
	case "rm":
		return false, cli.RM(args[1:])
	case "pwd":
		return false, cli.PWD()
	case "help":
		return false, cli.HELP()
	case "exit":
		return true, nil
	case "clear":
		fmt.Print("\033[H\033[2J")
		return false, nil
	default:
		return false, fmt.Errorf("fm_tools: Unknown command! (Type `help` for help!)")
	}
}

func parseInput(input string) []string {
	parts := strings.Fields(strings.ReplaceAll(input, `\ `, "\x00"))
	for i := range parts {
		parts[i] = strings.ReplaceAll(parts[i], "\x00", " ")
	}
	return parts
}
