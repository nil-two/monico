monico
======

Execute COMMAND for changes to current directory.

```
# execute make for changes to ./
monico make
```

Requirements
-----------

- bash
- inotify-tools

Installation
------------

monico is a simple shell script.

The following instructions assume that `~/bin` is on your `$PATH`.
If that is not the case, you can substitute your favorite location.

```sh
curl -L https://raw.githubusercontent.com/nil2nekoni2/monico/master/monico > ~/bin/monico
chmod 755 ~/bin/monico
```

Usage
-----

```
$ monico [OPTION]... COMMAND [ARGS]...
Execute COMMAND for changes to current directory.

Options:
  -b, --buffer-output    output after command finished
  -C, --no-clear         suppress clearing at execute
  -d, --directory=PATH   change the monitored directory to PATH
      --help             display this help text and exit
      --version          display version information and exit


# execute go test for changes to ./
monico go test

# execute make for changes to ./src
monico -d src make
```

License
-------

MIT License

Author
------

nil2 <nil2@nil2.org>
