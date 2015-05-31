monico
======

Run COMMAND if there is modified current directory

	$ monico make
	(Moniter current directory, and run 'make')

Usage
-----

	$ monico [OPTION]... COMMAND [ARGS]...
	Run COMMAND if there is modified current directory.

	Options:
		-d, --dir=PATH   moniter under PATH
		-h, --help       show this help message

Installation
------------

###compiled binary(not yet)

See [releases](https://github.com/kusabashira/monico/releases)

###go get

	go get github.com/kusabashira/lclip/cmd/lclip

Example
-------

	$ monico go test
	(Moniter current directory, and run 'go test')

	$ monico -d=src make
	(Moniter ./src , and run 'make')


License
-------

MIT License

Author
------

wara <kusabashira227@gmail.com>
