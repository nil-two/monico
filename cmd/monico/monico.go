package main

import (
	"os"
)

func usage() {
	os.Stderr.WriteString(`
Usage: monico [OPTION]... COMMAND [ARGS]...
Run COMMAND if there is modified current directory.

Options:
	-d, --dir=PATH   cd PATH before monitering
	-h, --help       show this help message
`[1:])
}

func main() {
}
