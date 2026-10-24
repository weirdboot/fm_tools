package cli

import "fmt"

var options []string = []string{
	" cp: cp `source` `destination`",
	" cd: cd `destination`",
	" rm: rm `file_dest_1` `file_dest_2` `file_dest_3` ...\n   -> Use \"-d\" flag to delete non-empty directory",
	" ls: ls `destination`",
	" grep: `source` `word1` `word2` `word3` ...",
	" mv: mv `source` `destination`",
	" pwd: pwd",
	" clear: clear",
	" exit: exit",
}

func HELP() error {
	fmt.Println("USAGE:")
	for i := range options {
		fmt.Println(options[i])
	}
	return nil
}
