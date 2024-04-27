# TinyTimer

TinyTimer is a small timer, written in Golang. It is meant to be a terminal
based timer for solving Rubix cubes or other twisty puzzles. Its main focus
is to be as sleek and as out of your way as possible -- which is why the
interface is just the timer, nothing else.

## Main features
- Splits
- Storing PBs
- Logging to file to log each solve

## Compiling

Install Go.
```
$ cd src/
$ mkdir -p bin/
$ go build main.go -o ../bin/tinytimer
```