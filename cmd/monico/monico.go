package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
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

func newCommand(a []string) *exec.Cmd {
	c := exec.Command(a[0], a[1:]...)
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}

func copyCommand(src *exec.Cmd) *exec.Cmd {
	return newCommand(src.Args)
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
