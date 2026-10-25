### About Project

This project, called `fm_tools` (`fm` is short of `file manager`), is intended to manage files from the command line, much like a shell.  
The goal of the project is to create a fully functional command-line file manager.

### Setup & Usage
#### From source
**Requirements:** `go1.22` and above
1. Clone the repository with Git
   ````shell
   git clone https://github.com/weirdboot/fm_tools.git
   cd fm\_tools
   go build -trimpath -ldflags="-s -w" -o app .```
   ````
2. Run program
   `./app`
#### Binary releases
Download the binary for your platform from [Releases](https://github.com/weirdboot/fm_tools/releases).

**Linux/macOS:**
```
chmod +x fm-linux-amd64
./fm-linux-amd64
```
**Windows:**
```
fm-windows-amd64.exe
```

### Supported CLI commands

| Command | Definition                                                                                     | Usage                                                        | Example                                                    |
| ------- | ---------------------------------------------------------------------------------------------- | ------------------------------------------------------------ | ---------------------------------------------------------- |
| cp      | copying file/directory from source to destination                                              | `cp "source" "destination"`                                  | `copy ./main.go ./go.main`                                 |
| cd      | goes to the specified directory                                                                | `cd "destination"`                                           | `cd /home/weirdboot`                                       |
| rm      | deletes the specified file or (empty) directory<br>Use `-d` flag to delete non-empty directory | `rm "file_dest" "empty_dir_dest"`<br>`rm -d "non_empty_dir"` | `rm -d ./non_empty_dir`<br>`rm ../photos2009 ../script.sh` |
| ls      | displays files in the specified directory                                                      | `ls "destination"`<br>`ls`                                   | `ls`<br>`ls /home/weirdboot /usr`                          |
| grep    | searches for the desired word in the file and prints the line                                  | `grep "source_file" "word_1" "word_2" "word_3" ...`          | `grep ./main.go func if package`                           |
| mv      | moving a file/directory to a designated location                                               | `mv "source" "destination"`                                  | `mv ./cool_folder ./folder_cool`                           |
| pwd     | shows current directory                                                                        | `pwd`                                                        | `pwd`                                                      |
| clear   | clears command line history                                                                    | `clear`                                                      | `clear`                                                    |
| help    | shows how to use commands                                                                      | `help`                                                       | `help`                                                     |
| exit    | completes the program                                                                          | `exit`                                                       | `exit`                                                     |

### Showcase

![Showcase](showcase.gif)
