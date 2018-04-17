# Buffalo Upgrade(x)

This plugin is a proposal for a new `buffalo upgrade` command that will update Buffalo, Pop, Soda, and the user's application. Flags can turn on/off the different parts and if the user runs it in a non-Buffalo application directory it will automatically skip that step.

## Installation

```bash
$ go get -u -v github.com/gobuffalo/buffalo-upgradex
```

## Usage

```bash
$ buffalo upgradex
```

### Flags

* `--sqlite`: enables SQlite3 support in the binaries and packages
* `--skip-buffalo`: skips updating the `buffalo` binary
* `--skip-pop`: skips updating the `soda` binary
* `--skip-app`: skips updating the current application (if inside a Buffalo application)
