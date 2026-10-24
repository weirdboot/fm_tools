package cli

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

type found struct {
	row_num int
	word    string
	line    string
}

func GREP(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("grep: expected atleast 2 arguments, but %d were received", len(args))
	}
	read, err := os.ReadFile(args[0])
	if err != nil {
		return fmt.Errorf("grep: %w", err)
	}
	lines := strings.Split(string(read), "\n")
	args = args[1:]

	var result []found = []found{}
	for i := range args {
		for j := range lines {
			if strings.Contains(lines[j], args[i]) {
				result = append(result, found{
					row_num: j,
					word:    args[i],
					line:    strings.TrimLeftFunc(lines[j], unicode.IsSpace),
				})
			}
		}
	}

	for i := range result {
		fmt.Printf("%d -> %s || %s\n", result[i].row_num, result[i].word, result[i].line)
	}
	return nil
}
