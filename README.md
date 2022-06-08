# マネ (mane)

[mane](github.com/shimst3r/mane) is a currency locale formatting tool.

## Examples

It works from `$STDIN` to `$STDOUT`

```shell
$ mane -s EU -t US 1.000,00
1,000.00
```

and from `$FILE` to `$FILE`

```shell
$ mane -s EU -t US -f deutsch.txt -o english.txt
Formatted 100 lines
```

as well as any combination

```shell
$ mane -s EU -t US -f deutsch.txt
1,000.00
-9,15
450.95
```

## Installation

mane is a Go tool, so install it via the `go` command:


``shell
# go version < 1.16
$ go get github.com/shimst3r/mane

# go version >= 1.16
$ go install github.com/shimst3r/mane
```
