package main

import (
	"flag"
	"fmt"
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
	var moniterPath string
	flag.StringVar(&moniterPath, "d", "", "")
	flag.StringVar(&moniterPath, "dir", "", "")

	var isHelp bool
	flag.BoolVar(&isHelp, "h", false, "")
	flag.BoolVar(&isHelp, "help", false, "")
	flag.Usage = shortUsage
	flag.Parse()
	switch {
	case isHelp:
		usage()
		os.Exit(0)
	case flag.NArg() < 1:
		fmt.Fprintln(os.Stderr, "monico:", "no specify COMMAND")
		shortUsage()
		os.Exit(2)
	}
}
