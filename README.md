monico
======

Execute COMMAND for changes to current directory.

```
$ monico make
(execute make for changes to current directory)
```

Requirement
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


$ monico go test
(execute go test for changes to current directory)

$ monico -d=src make
(execute make for changes to ./src)
```

License
-------

MIT License

Author
------

wara <kusabashira227@gmail.com>
