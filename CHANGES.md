### v0.3.0 - 2015-11-20

- Support `-C`, `--no-clear` to allow suppressing clearing at execute.
- Remove short option for `--help`.
- Remove short option for `--version`.
- Change the format of version from "v0.3.1" to "0.3.1".

### v0.2.0 - 2015-11-06

- Stop flag parsing at `--`.
- Interpret equal sign after `-d` as directory's name.
  - later:   "-d=src" #=> DIRECTORY = "=src"
  - earlier: "-d=src" #=> DIRECTORY = "src"

### v0.1.1 - 2015-09-09

- Read argument after `--directory` as target directory.
- Allow argument which includes spaces.

### v0.1.0 - 2015-08-09

- Initial release.
