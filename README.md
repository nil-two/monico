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
- inotify

Installation
------------

monico is a simple shell script.

The following instructions assume that `~/bin` is on your `$PATH`.
If that is not the case, you can substitute your favorite location.

```sh
curl -L https://raw.githubusercontent.com/kusabashira/monico/master/monico > ~/bin/monico
chmod 755 ~/bin/monico
```

Usage
-----

```
$ monico [OPTION]... COMMAND [ARGS]...

Options:
  -d, --directory=PATH   Change monitored directory to PATH
  -h, --help             display this help text and exit
  -v, --version          display version information and exit


# execute go test for changes to ./
monico go test

# execute make for changes to ./src
monico -d=src make
```

License
-------

MIT License

Author
------

kusabashira <kusabashira227@gmail.com>
