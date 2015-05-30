package main

import (
	"os"
)

func shortUsage() {
	os.Stderr.WriteString(`
Usage: monico [OPTION]... COMMAND [ARGS]...
Try 'grep --help' for more information.
`[1:])
}

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
