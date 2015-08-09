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

wara <kusabashira227@gmail.com>
